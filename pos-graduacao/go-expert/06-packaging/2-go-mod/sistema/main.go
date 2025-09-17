package main

import (
	"fmt"
	// um modo de importar o pacote é usando o caminho relativo, porém n é recomendado
	// go mod edit -replace (...)=../math
	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/06_packaging/2-go-mod/math"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m)
}
