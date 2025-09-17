package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	// Muitas vezes, quando temos uma api, precisamos estabelecer limites, como o timeout
	client := http.Client{
		Timeout: time.Second, // 1 segundo
	}

	// resp, err := client.Get("https://www.google.com")
	jsonVar := bytes.NewBuffer([]byte(`{"nome":"Bruno"}`)) // precisamos bufferizar o json para ele poder ser enviado
	resp, err := client.Post("https://www.google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(body))
	io.CopyBuffer(os.Stdout, resp.Body, nil)

	// Podemos configurar uma request da seguinte forma também
	req2, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		panic(err)
	}
	req2.Header.Add("Aceppt", "application/json")
	resp2, err := client.Do(req2)
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()
	io.CopyBuffer(os.Stdout, resp2.Body, nil)
}
