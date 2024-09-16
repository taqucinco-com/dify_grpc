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

	"github.com/joho/godotenv"
	sse "github.com/tmaxmax/go-sse"
)

type EventData struct {
	Event  string  `json:"event"`
	TaskID string  `json:"task_id"`
	Answer *string `json:"answer"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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
	difyReq, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.dify.ai/v1/chat-messages", bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
		return
	}
	difyReq.Header.Set("Authorization", "Bearer "+key)
	difyReq.Header.Set("Content-Type", "application/json")

	conn := sseClient.NewConnection(difyReq)

	unsubscribe := conn.SubscribeMessages(func(ev sse.Event) {
		// fmt.Printf(ev.Data + "\n")
		var eventData EventData
		err := json.Unmarshal([]byte(ev.Data), &eventData)
		if err == nil {
			eventName := eventData.Event
			taskID := eventData.TaskID

			if eventName == "workflow_started" {
				fmt.Println("Task ID:", taskID)
			} else if eventName == "message" {
				answer := *eventData.Answer
				fmt.Println("Answer:", answer)
			}
		}

	})
	defer unsubscribe()

	if err := conn.Connect(); !errors.Is(err, io.EOF) {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
