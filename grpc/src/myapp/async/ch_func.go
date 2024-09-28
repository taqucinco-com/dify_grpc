package main

import (
	"fmt"
	"sync"
	// "sync"
)

type Result struct {
	Error   error
	Message string
}

func send(wg *sync.WaitGroup) <-chan Result {
	defer wg.Done()
	ch := make(chan Result)
	defer close(ch)
	ch <- Result{Error: nil, Message: "Hello"}
	return ch
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := send(&wg)
	wg.Wait()
	for {
		message, ok := <-ch
		if !ok {
			fmt.Println("channel closed")
			break
		}
		fmt.Println("[send]:", message)
	}
}
