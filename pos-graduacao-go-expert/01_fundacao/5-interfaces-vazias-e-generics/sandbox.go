package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Hello, World!"

	showType(x)
	showType(y)

	var minhaVar interface{} = "Bruno Reinoso"
	println(minhaVar.(string))
	res, ok := minhaVar.(int) // sem o ok neste caso vai gerar panic, pq n vai funcionar (type assertion)
	fmt.Printf("O valor de res é %v e ok é %v\n", res, ok)

	// ========================================================

	m := map[string]int{"Bruno": 2000, "Aurelio": 80000}
	m2 := map[string]float64{"Bruno": 2000.0, "Aurelio": 80000.1}
	m3 := map[string]MyNumber{"Bruno": 2000, "Aurelio": 8002}

	fmt.Println(Soma(m))
	fmt.Println(Soma2(m2))
	fmt.Println(Soma2(m3))

}

type MyNumber int

// Constraint - https://pkg.go.dev/golang.org/x/exp/constraints
type Number interface {
	~int | float64 // o til é para o go entender que um outro tipo, como MyNumber,
	// pode ser identificado como Number no caso de ser int também
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func Soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

func Soma2[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é %v\n", t, t)
}
