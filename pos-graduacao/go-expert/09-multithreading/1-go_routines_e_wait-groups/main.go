package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is working...\n",i, name)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1
func main() {
	waitGroup := &sync.WaitGroup{}
	// Adiciona 3 "créditos" ao WaitGroup
	// Isso significa que o WaitGroup vai esperar as 3 goroutines terminarem
	waitGroup.Add(3)

	// Thread 2
	go task("A", waitGroup)
	// Thread 3
	go task("B", waitGroup)
	// Thread 4
	go func() {
		defer waitGroup.Done()
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is working...\n",i, "Anonymous")
			time.Sleep(1 * time.Second)
		}
	}()

	waitGroup.Wait()
}
