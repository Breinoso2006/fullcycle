package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Criação do arquivo

	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// Escrita no arquivo

	// se eu souber que o dado que quero escrever é string
	// tamanho, err := f.WriteString("Hello, World!")

	tamanho, err := f.Write([]byte("Escrevendo dados no arquivo!"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho de %d bytes\n", tamanho)
	f.Close()

	// Leitura do arquivo

	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	// Preciso converter pois o que está no arquivo são bytes
	fmt.Println((string(arquivo)))

	// Leitura de arquivos muito grandes
	// Leitura em partes, utilizando buffer
	// buffer é uma área de memória intermediária usada para armazenar dados temporariamente

	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 4)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Print(string(buffer[:n]) + "--")
	}

	arquivo2.Close()

	// Remover arquivo

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
