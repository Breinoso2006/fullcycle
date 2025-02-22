package main

func main(){
	// no go temos apenas o for, com 4 formas de uso

	// for com uma condição, iterando sobre i
	for i := 0; i < 10; i++ {
		println(i)
	}

	// for com range (lembrando que posso usar blank identifier _)
	numeros := []int{1, 2, 3, 4, 5}
	for indice, valor := range numeros {
		println(indice, valor)
	}

	// for como while
	contador := 0
	for contador < 10 {
		println(contador)
		contador++
	}

	// for infinito
	for {
		println("executando infinitamente")
		break
	}

	// Condicionais

	// if/else (não temos else if ou algo do tipo)
	// && (and) e || (or)
	i := 3
	if i == 1 {
		println("1 é igual a 1")
	} else {
		println("1 é diferente de 1")
	}

	// switch
	// não é necessário usar break, ele já faz isso por padrão
	// posso usar uma expressão no case
	switch i {
	case 1:
		println("i é igual a 1")
	case 2:
		println("i é igual a 2")
	default:
		println("i é diferente de 1 e 2")
	}
}