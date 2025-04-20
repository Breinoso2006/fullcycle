package main

import (
	"fmt"

	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/06_packaging/1-intro/math"
)

func main() {
	m := math.Math{
		A: 10,
		B: 20,
	}
	s := m.Sum()
	fmt.Println(s)
}
