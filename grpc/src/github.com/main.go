package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	sse "github.com/tmaxmax/go-sse"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	key := os.Getenv("DIFY_API_KEY")
	fmt.Println(key)

	client := sse.Client{
		Backoff: sse.Backoff{
			MaxRetries: -1,
		},
	}
	json := `{
	    "inputs": {},
    	"query": "What are the specs of the iPhone 13 Pro Max?",                  
    	"response_mode": "streaming",
    	"conversation_id": "",
    	"user": "user-123",
    	"files": []
	}`
	req, err := http.NewRequest(http.MethodPost, "https://api.dify.ai/v1/chat-messages", bytes.NewBuffer([]byte(json)))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+key)
	req.Header.Set("Content-Type", "application/json")

	conn := client.NewConnection(req)

	unsubscribe := conn.SubscribeMessages(func(ev sse.Event) {
		fmt.Printf(ev.Data + "\n")
	})
	defer unsubscribe()

	if err := conn.Connect(); !errors.Is(err, io.EOF) {
		fmt.Fprintln(os.Stderr, err)
	}
}
