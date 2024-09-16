package main

import (
	"fmt"
	"net/http"
	"time"
)

func hoge() {
	http.HandleFunc("/sse", sseHandler)
	http.ListenAndServe(":8081", nil)
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/sse Handler called")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	clientGone := r.Context().Done()
	rc := http.NewResponseController(w)
	t := time.NewTicker(time.Second)
	defer t.Stop()

	for i := 0; i < 5; i++ {
		select {
		case <-clientGone:
			fmt.Println("Client disconnected")
			return
		case <-t.C:
			fmt.Fprintf(w, "data: %s\n\n", time.Now().Format(time.Stamp))
			fmt.Println("Sent event:", i)
			rc.Flush()
		}
	}
	fmt.Println("sse finish")
}
