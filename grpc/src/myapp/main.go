package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"taqucinco.com/myapp/dify"

	"github.com/joho/godotenv"
	sse "github.com/tmaxmax/go-sse"
)

type Result struct {
	Error   error
	Message *string
}

func send(ch chan<- Result /*, wg *sync.WaitGroup*/) {
	// defer wg.Done()
	defer close(ch)

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
		ch <- Result{Error: envErr, Message: nil}
		return
	}
	key := os.Getenv("DIFY_API_KEY")
	// fmt.Println(key)

	sseClient := sse.Client{
		Backoff: sse.Backoff{
			MaxRetries: -1,
		},
	}
	userId := "user-123"
	body := `{
		"inputs": {},
		"query": "What are the specs of the iPhone 13 Pro Max?",                  
		"response_mode": "streaming",
		"conversation_id": "",
		"user": "` + userId + `",
		"files": []
	}`
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	difyReq, sseErr := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.dify.ai/v1/chat-messages", bytes.NewBuffer([]byte(body)))
	if sseErr != nil {
		log.Fatalf("Error creating request: %v", sseErr)
		ch <- Result{Error: sseErr, Message: nil}
		return
	}
	difyReq.Header.Set("Authorization", "Bearer "+key)
	difyReq.Header.Set("Content-Type", "application/json")

	conn := sseClient.NewConnection(difyReq)

	unsubscribe := conn.SubscribeMessages(func(ev sse.Event) {
		fmt.Printf(ev.Data + "\n")
		var eventData dify.EventData
		err := json.Unmarshal([]byte(ev.Data), &eventData)
		if err == nil {
			eventName := eventData.Event
			taskID := eventData.TaskID

			if eventName == "workflow_started" {
				message := `{"task_id":"` + taskID + `"}`
				ch <- Result{Error: nil, Message: &message}
			} else if eventName == "message" {
				answer := `{"answer":"` + *eventData.Answer + `"}`
				ch <- Result{Error: nil, Message: &answer}
			}
		} else {
			ch <- Result{Error: err, Message: nil}
		}
	})
	defer unsubscribe()

	err := conn.Connect()
	if errors.Is(err, io.EOF) {
		fmt.Fprintln(os.Stderr, err)
	} else if err != nil {
		log.Fatalf("err: %v", err)
		ch <- Result{Error: err, Message: nil}
	}
}

func main() {
	ch := make(chan Result)
	go send(ch)
	for {
		result, ok := <-ch
		if !ok {
			fmt.Println("channel closed")
			break
		}
		if result.Message != nil {
			log.Println("Message:", *result.Message)
		} else if result.Error != nil {
			log.Println("Error:", result.Error)
		}
	}
}
