package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DolarFromApi struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {
	http.HandleFunc("/", HealthCheck)
	http.HandleFunc("/cotacao", DolarQuoteNow)
	http.ListenAndServe(":8080", nil)
}

type DolarQuote struct {
	ID    uint `gorm:"primaryKey"`
	Price string
}

func DolarQuoteNow(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
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

	var data DolarFromApi
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshalling JSON: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{
		"bid": data.Usdbrl.Bid,
	}
	json.NewEncoder(w).Encode(response)

	RegisterQuoteInDatabase(data.Usdbrl.Bid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error registering quote in database: %v\n", err)
	}
}

func RegisterQuoteInDatabase(bid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	db, err := gorm.Open(sqlite.Open("meubanco.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	err = db.AutoMigrate(&DolarQuote{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	quote := DolarQuote{Price: bid}
	err = db.WithContext(ctx).Create(&quote).Error
	if err != nil {
		return fmt.Errorf("failed to create quote: %w", err)
	}
	return nil
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Healthy"))
	w.WriteHeader(http.StatusOK)
}
