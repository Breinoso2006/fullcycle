package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	allocateMemory := func(size int) []byte {
		return make([]byte, size)
	}
	for i := 0; i < 10; i++ {
		allocateMemory(1024 * 1024 * 10)
		time.Sleep(1 * time.Second)
	}
	// exibindo uso de memória
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc: %v MB\n", m.Alloc/1024/1024)
	fmt.Printf("TotalAlloc: %v MB\n", m.TotalAlloc/1024/1024)
	fmt.Printf("Sys: %v MB\n", m.Sys/1024/1024)
	fmt.Printf("NumGC: %v\n", m.NumGC)

}
