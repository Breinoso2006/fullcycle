package main

import (
	"fmt"
	"os"
)

func main() {
	i := 0
	// Create many files in a loop
	for i < 100 {
		f, err := os.Create(fmt.Sprintf("./tmp/file-%d.txt", i))
		if err != nil {
			panic(err)
		}
		f.WriteString("Hello, World!")
		i++
	}
}
