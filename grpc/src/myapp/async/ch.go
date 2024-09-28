package main

import (
	"fmt"
	// "sync"
	"time"
)

func send(ch chan<- string /*, wg *sync.WaitGroup*/) {
	// defer wg.Done()
	defer close(ch)
	ch <- "Hello"
	time.Sleep(time.Second)
	ch <- " World."
}

func main() {
	// var wg sync.WaitGroup
	// wg.Add(1)
	ch := make(chan string)
	go send(ch)
	for {
		message, ok := <-ch
		if !ok {
			fmt.Println("channel closed")
			break
		}
		fmt.Println("[send]:", message)
	}
	// wg.Wait()
}
