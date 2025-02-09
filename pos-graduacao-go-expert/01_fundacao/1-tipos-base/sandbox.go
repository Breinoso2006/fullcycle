package main

import (
	"fmt"
)

// const a = "Hello, World!"

// type ID int

// var (
// 	b bool    = true
// 	c int     = 10
// 	d string  = "Teste"
// 	e float64 = 13.5
// 	f ID      = 12345
// )

func main() {
	// 	println("Testando Go!")
	// 	h := "Hello, World!"
	// 	println(h)

	// 	var meuAarray [4]int
	// 	meuAarray[0] = 10
	// 	meuAarray[1] = 20
	// 	meuAarray[2] = 30

	// 	fmt.Printf("O tamanho do meu array é %d\n", len(meuAarray))
	// 	fmt.Printf("O ultimo elemento do meu array é %d\n", meuAarray[len(meuAarray)-1])

	// 	for index, value := range meuAarray {
	// 		fmt.Printf("O valor do elemento %d é %d\n", index, value)
	// 	}

	// 	s := []int{1, 2, 3, 4, 5}

	// 	fmt.Printf("O tamanho do meu slice é %d, capacidade %d e tem os números %v\n", len(s), cap(s), s)
	// 	fmt.Printf("O tamanho do meu slice é %d, capacidade %d e tem os números %v\n", len(s[:2]), cap(s[:2]), s[:2])
	// 	fmt.Printf("O tamanho do meu slice é %d, capacidade %d e tem os números %v\n", len(s[2:]), cap(s[2:]), s[2:])
	// 	fmt.Printf("O tamanho do meu slice é %d, capacidade %d e tem os números %v\n", len(s[1:1]), cap(s[1:1]), s[1:1])
	// 	fmt.Printf("O tamanho do meu slice é %d, capacidade %d e tem os números %v\n", len(s[1:3]), cap(s[1:3]), s[1:3])
	// 	s = append(s, 6)
	// 	fmt.Printf("Após o append o tamanho do meu slice é %d, capacidade %d e tem os números %v\n", len(s), cap(s), s)

	salarios := map[string]int{
		"José":  1000,
		"Maria": 5000,
		"Pedro": 2000,
		"Bruno": 3000,
	}
	delete(salarios, "Pedro")
	salarios["Mauricio"] = 4500
	fmt.Println(salarios)

	// iniciar sem valores
	// sal := make(map[string]int)
	// sal2 := map[string]int{}

	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}
}
