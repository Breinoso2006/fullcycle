package main

import (
	"sync"
	"time"
)

func main() {
	channel := make(chan int)
	numberOfWorkers := 10
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := range numberOfWorkers {
		go worker(i, channel, &wg)
	}

	for i := range 100 {
		channel <- i
	}
	wg.Wait()
	close(channel)
}

func worker(idWorker int, data chan int, wg *sync.WaitGroup) {
	for d := range data {
		println("Worker", idWorker, "processing data", d)
		time.Sleep(time.Second)
		wg.Done()
	}
}