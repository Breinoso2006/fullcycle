package main

import (
	"breinoso2006/labs-go/otel/server/internal/config"
	"breinoso2006/labs-go/otel/server/internal/handler"
	"breinoso2006/labs-go/otel/server/internal/telemetry"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		// Arquivo .env não encontrado, usar valores padrão
		cfg = &config.Config{
			Port:            "8080",
			OrchestratorURL: "http://localhost:8081",
			ServiceName:     "service-a",
			ZipkinURL:       "http://localhost:9411/api/v2/spans",
		}
	}

	// Sobrescrever com variáveis de ambiente se existirem
	if envPort := os.Getenv("PORT"); envPort != "" {
		cfg.Port = envPort
	}

	if envOrchURL := os.Getenv("ORCHESTRATOR_URL"); envOrchURL != "" {
		cfg.OrchestratorURL = envOrchURL
	}

	if envServiceName := os.Getenv("SERVICE_NAME"); envServiceName != "" {
		cfg.ServiceName = envServiceName
	}

	if envZipkin := os.Getenv("ZIPKIN_ENDPOINT"); envZipkin != "" {
		cfg.ZipkinURL = envZipkin
	}

	handler.SetOrchestratorURL(cfg.OrchestratorURL)

	shutdown, err := telemetry.InitTracer(
		cfg.ServiceName,
		cfg.ZipkinURL,
	)
	if err != nil {
		log.Fatalf("Erro ao inicializar tracer: %v", err)
	}
	defer func() {
		if err := shutdown(context.Background()); err != nil {
			log.Printf("Erro ao encerrar tracer: %v", err)
		}
	}()

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Servidor iniciando na porta %s", cfg.Port)
	log.Printf("Service Name: %s", cfg.ServiceName)
	log.Printf("Orchestrator URL: %s", cfg.OrchestratorURL)
	log.Printf("Zipkin Endpoint: %s", cfg.ZipkinURL)

	http.HandleFunc("/weather", handler.WeatherHandler)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
