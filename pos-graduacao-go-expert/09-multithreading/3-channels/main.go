package main

import "sync"

func main() {
	channel := make(chan string)

	go func() {
		// Preenche o canal com uma string
		channel <- "Hello, world!"
	}()

	// Recebe a string do canal
	msg := <-channel
	println(msg)

	// For com canais
	channel2 := make(chan int)

	go publish(channel2)
	reader(channel2)

	// For com canais e sync.WaitGroup
	wg := sync.WaitGroup{}
	wg.Add(10)
	channel3 := make(chan int)

	go publish(channel3)
	go reader2(channel3, &wg)
	wg.Wait()

}

func reader(ch chan int) {
	for x := range ch {
		println(x)
	}
}

func reader2(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		println(x)
		wg.Done()
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	// se não fechar o canal teremos um deadlock
	close(ch)
}