package main

import (
	"breinoso2006/desafio-cep/internal/config"
	"breinoso2006/desafio-cep/internal/handler"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Printf("AVISO: Erro ao carregar configurações: %v", err)
		log.Printf("Tentando iniciar servidor mesmo assim na porta %s...", port)
	}

	cepHandler := handler.NewCepHandler(cfg.WeatherAPIKey)

	http.HandleFunc("/weather/", cepHandler.GetTemperatureByCep)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Servidor iniciando na porta %s", port)
	log.Printf("WeatherAPIKey configurada: %v", cfg.WeatherAPIKey != "")

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
