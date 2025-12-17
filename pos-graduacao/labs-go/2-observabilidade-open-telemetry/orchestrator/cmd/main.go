package main

import (
	"breinoso2006/labs-go/otel/orchestrator/internal/config"
	"breinoso2006/labs-go/otel/orchestrator/internal/handler"
	"breinoso2006/labs-go/otel/orchestrator/internal/telemetry"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Printf("AVISO: Erro ao carregar configurações: %v", err)
		log.Println("Usando valores padrão")
		cfg = &config.Config{
			Port:           "8081",
			ServiceName:    "service-b",
			ZipkinEndpoint: "http://localhost:9411/api/v2/spans",
		}
	}

	if cfg.Port == "" {
		cfg.Port = os.Getenv("PORT")
		if cfg.Port == "" {
			cfg.Port = "8081"
		}
	}

	if cfg.ServiceName == "" {
		cfg.ServiceName = os.Getenv("SERVICE_NAME")
		if cfg.ServiceName == "" {
			cfg.ServiceName = "service-b"
		}
	}

	if cfg.ZipkinEndpoint == "" {
		cfg.ZipkinEndpoint = os.Getenv("ZIPKIN_ENDPOINT")
		if cfg.ZipkinEndpoint == "" {
			cfg.ZipkinEndpoint = "http://localhost:9411/api/v2/spans"
		}
	}

	// Inicializar OTEL
	shutdown, err := telemetry.InitTracer(cfg.ServiceName, cfg.ZipkinEndpoint)
	if err != nil {
		log.Fatalf("Erro ao inicializar tracer: %v", err)
	}
	defer func() {
		if err := shutdown(context.Background()); err != nil {
			log.Printf("Erro ao encerrar tracer: %v", err)
		}
	}()

	cepHandler := handler.NewCepHandler(cfg.WeatherAPIKey)

	http.HandleFunc("/weather/", cepHandler.WeatherHandler)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Servidor iniciando na porta %s", cfg.Port)
	log.Printf("Service Name: %s", cfg.ServiceName)
	log.Printf("Zipkin Endpoint: %s", cfg.ZipkinEndpoint)
	log.Printf("WeatherAPIKey configurada: %v", cfg.WeatherAPIKey != "")

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
