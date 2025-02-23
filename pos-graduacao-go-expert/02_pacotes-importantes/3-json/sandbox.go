package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 1000}

	// Struct para json, guardando o valor
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	print(string(res))

	// Já faz a transformação da struct para json e envia para o stdout
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	// Utilizando o que foi configurado nas tags
	jsonEx := []byte(`{"n":1,"s":1000}`)

	// variável vazia
	var contaX Conta

	// modificamos o valor da variável no endereço de memória
	err = json.Unmarshal(jsonEx, &contaX)
	println(contaX.Saldo)
}
