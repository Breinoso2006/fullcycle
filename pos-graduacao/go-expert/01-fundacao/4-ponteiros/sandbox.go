package main

func main() {
	// testes()
	// 	var (
	// 		a int = 10
	// 		b int = 20
	// 	)

	// 	println(soma(a, b))
	// 	println(somaETrocaValores(&a, &b))
	// 	println(a, b)

	var poupanca Conta = Conta{
		saldo: 1000,
	}

	println(poupanca.Deposito(200))
	println(poupanca.saldo)

	var corrente Conta = *NewConta()
	println(corrente.saldo)
}

type Conta struct {
	saldo int
}

func (c *Conta) Deposito(valor int) int {
	c.saldo += valor
	return c.saldo
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func soma(a, b int) int {
	return a + b
}

func somaETrocaValores(a, b *int) int {
	c := *a
	*a = *b
	*b = c
	return *a + *b
}

func testes() {
	var a int = 10
	var ponteiro *int = &a // ponteiro para a variável a

	println(a)  // imprime o valor da variável a
	println(&a) // imprime o endereço de memória da variável a

	println(ponteiro)  // imprime o endereço de memória da variável a
	println(*ponteiro) // imprime o valor da variável a através do ponteiro

	*ponteiro = 20 // altera o valor da variável a através do ponteiro
	println(a)     // imprime o valor da variável a

	b := &a

	println(b)  // imprime o endereço de memória da variável a
	println(*b) // imprime o valor da variável a através do ponteiro

	var ponteiro2 **int = &b // ponteiro para o ponteiro da variável a
	**ponteiro2 = 50
	println(a) // imprime o valor da variável a

	var ponteiro3 ***int = &ponteiro2 // ponteiro para o ponteiro do ponteiro da variável a
	***ponteiro3 = 100
	println(a)          // imprime o valor da variável a
	println(&ponteiro3) // imprime o endereço de memória do ponteiro3
	// (que é o endereço de memória do ponteiro2, que é o endereço
	// de memória do ponteiro, que é o endereço de memória da variável a)
}
