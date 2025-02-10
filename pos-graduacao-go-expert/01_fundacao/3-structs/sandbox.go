package main

import "fmt"

// Interfaces só podem implementar métodos, não atributos
type Animal interface {
	Andar()
	// Saltar() teria que ser implementado por todos os tipos que implementam a interface Animal
}

func FormaDeAndar(animal Animal) {
	fmt.Print("Como eu ando?\n")
	animal.Andar()
}

type Pessoa struct {
	Nome   string
	Idade  int
	Altura float64
}

type Funcionario struct {
	Pessoa
	Salario float64
	Ativo   bool
}

func (f Funcionario) promover() float64 {
	f.Salario += 3000
	return f.Salario
}

func (p Pessoa) Andar() {
	fmt.Println("Sou uma pessoa, e ando descalço.")
}

func (f Funcionario) Andar() {
	fmt.Println("Sou um funcionário, ando de sapato.")
}

func main() {
	bruno := Pessoa{
		Nome:   "Bruno",
		Idade:  28,
		Altura: 1.62,
	}
	// fmt.Println(bruno)

	// no go utilizamos composição ao invés de herança
	aurelio := Funcionario{
		Pessoa: Pessoa{
			Nome:   "Aurelio",
			Idade:  30,
			Altura: 1.75,
		},
		Salario: 5000.0,
		Ativo:   true,
	}

	// fmt.Println(aurelio)
	// aurelio.promover()
	// fmt.Println(aurelio)
	// o método promover não altera o valor do salário caso não seja atribuído a uma variável
	// aurelio.Salario = aurelio.promover()
	// fmt.Println(aurelio)
	// aurelio.Andar()

	// Ambos implementam a interface Animal
	FormaDeAndar(aurelio)
	FormaDeAndar(bruno)
}
