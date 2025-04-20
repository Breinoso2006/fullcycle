package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Dolar struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Printf("Request timed out after 300ms.\n")
			return
		}
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	var dolar Dolar
	err = json.Unmarshal(body, &dolar)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return
	}

	err = RegisterQuoteInTxt(dolar.Bid)
	if err != nil {
		fmt.Printf("Error registering quote in txt: %v\n", err)
		return
	}
}

func RegisterQuoteInTxt(bid string) error {
	file, err := os.OpenFile("cotacao.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dolar: %s\n", bid))
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return err
	}

	fmt.Printf("Dolar: %v\n", bid)
	return nil
}
