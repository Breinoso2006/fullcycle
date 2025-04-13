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
		fmt.Fprintf(os.Stderr, "Error creating request: %v\n", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error making request: %v\n", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading response body: %v\n", err)
	}

	var dolar Dolar
	err = json.Unmarshal(body, &dolar)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshalling JSON: %v\n", err)
	}

	err = RegisterQuoteInTxt(dolar.Bid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error registering quote in txt: %v\n", err)
		return
	}
}

func RegisterQuoteInTxt(bid string) error {
	file, err := os.OpenFile("cotacao.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("Error opening/creating file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dolar: %s\n", bid))

	return err
}
