package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

func main() {
	// contexto vazio
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second) // cancela a execução após 1 segundo
	defer cancel()                                       // boa prática: chamar cancel após o uso do contexto

	req, err := http.NewRequestWithContext(ctx, "GET", "https://www.google.com", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))

}
