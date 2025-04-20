package main

import (
	"fmt"

	// a dependência é importada pelo workspace
	// go work init com os diretórios
	
	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/06_packaging/3-go-mod-rep/math"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m)
}
