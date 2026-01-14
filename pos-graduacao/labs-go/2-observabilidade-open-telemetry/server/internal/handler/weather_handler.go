package handler

import (
	"breinoso2006/labs-go/otel/server/internal/client"
	"encoding/json"
	"io"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type InputRequest struct {
	Cep string `json:"cep"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func SetOrchestratorURL(url string) {
	client.SetOrchestratorURL(url)
}

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tracer := otel.Tracer("service-a")
	ctx, span := tracer.Start(ctx, "POST /weather")
	defer span.End()

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Ler body JSON
	var input InputRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		span.SetAttributes(attribute.String("error", "invalid json"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "invalid zipcode"})
		return
	}

	cep := input.Cep
	span.SetAttributes(attribute.String("cep", cep))

	if !isValidCepFormat(cep) {
		span.SetAttributes(attribute.String("error", "invalid zipcode"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "invalid zipcode"})
		return
	}

	// Chama o orchestrator
	resp, err := client.WeatherOrchestratorRequest(ctx, cep)
	if err != nil {
		span.SetAttributes(attribute.String("error", err.Error()))
		span.SetStatus(codes.Error, "Failed to connect to orchestrator")
		http.Error(w, "Failed to connect to orchestrator", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copia os headers da resposta
	w.Header().Set("Content-Type", "application/json")

	// Copia o status code
	w.WriteHeader(resp.StatusCode)

	// Copia o body da resposta
	io.Copy(w, resp.Body)
}

func isValidCepFormat(cep string) bool {
	if len(cep) != 8 {
		return false
	}
	for _, c := range cep {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
