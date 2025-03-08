package main

import (
	"io"
	"net/http"
)

func main() {
	i := 0
	defer println("Testando o defer, i:", i)

	req, err := http.Get("https://www.google.com.br")
	if err != nil {
		panic(err)
	}
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
	// Precisamos fechar o corpo da requisição
	req.Body.Close()

	i = 1
	defer println("Testando o defer, i:", i)
	
}
