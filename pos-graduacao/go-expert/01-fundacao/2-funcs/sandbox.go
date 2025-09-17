package main

import (
	"fmt"
)

func main() {
	// fmt.Println(soma(1, 2))
	// valor, err := soma(20, 50)
	// fmt.Println(valor, err)
	fmt.Println(soma2(1, 2, 3, 4, 5))

	// closure
	result := func() int {
		fmt.Printf("%d\n", 2)
		return 2
	}()
	fmt.Println(result)

}

// func soma(a, b int) (int, error) { // pode ser também algo como func soma(a int, b int) int
// 	if a+b >= 50 {
// 		return a + b, errors.New("Resultado maior que 50")
// 	}
// 	return a + b, nil
// }

func soma2(numeros ...int) int {
	resultado := 0
	for _, numero := range numeros {
		resultado += numero
	}
	return resultado
}
