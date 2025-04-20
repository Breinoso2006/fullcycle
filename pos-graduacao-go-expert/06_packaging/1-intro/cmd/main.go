package main

import (
	"fmt"

	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/06_packaging/1-intro/math"
)

func main() {
	m := math.NewMath(1, 2)
	s := m.Add()
	fmt.Println(s)
	// apesar de conseguir printar m, não consigo acessar seus atributos
	fmt.Println(m)
}
