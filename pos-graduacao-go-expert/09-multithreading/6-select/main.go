package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	Id      int64
	Message string
}

func main() {
	channel := make(chan Message)
	channel2 := make(chan Message)
	var i int64 = 0

	go func() {
		for {
			time.Sleep(time.Second)
			atomic.AddInt64(&i, 1)
			message := Message{
				Id:      i,
				Message: "Hello from channel 1",
			}
			channel <- message
		}
	}()
	go func() {
		for {
			time.Sleep(2 * time.Second)
			atomic.AddInt64(&i, 1)
			message := Message{
				Id:      i,
				Message: "Hello from channel 2",
			}
			channel2 <- message
		}
	}()

	for {
		select {
		case value := <-channel:
			fmt.Printf("ID %d - %s\n ", value.Id, value.Message)
		case value := <-channel2:
			fmt.Printf("ID %d - %s\n ", value.Id, value.Message)
		case <-time.After(3 * time.Second):
			fmt.Printf("Timeout\n")
			// default:
			// 	fmt.Printf("No channel ready\n")
		}
	}
}
