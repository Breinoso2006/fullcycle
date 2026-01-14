package handler

import (
	"breinoso2006/labs-go/otel/orchestrator/internal/service"
	"encoding/json"
	"net/http"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type CepHandler struct {
	weatherAPIKey string
}

func NewCepHandler(weatherAPIKey string) *CepHandler {
	return &CepHandler{
		weatherAPIKey: weatherAPIKey,
	}
}

type TemperatureResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func (h *CepHandler) WeatherHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tracer := otel.Tracer("service-b")
	ctx, span := tracer.Start(ctx, "GET /weather/:cep")
	defer span.End()

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/weather/")
	cep := strings.TrimSpace(path)

	span.SetAttributes(attribute.String("cep", cep))

	infoCep, err := service.CepInfo(ctx, cep)
	if err != nil {
		span.SetAttributes(attribute.String("error", err.Error()))
		span.SetStatus(codes.Error, "Failed to fetch CEP info")
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(err.Error(), "CEP não encontrado") {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrorResponse{Message: "can not find zipcode"})
		} else {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(ErrorResponse{Message: "can not find zipcode"})
		}
		return
	}

	span.SetAttributes(
		attribute.String("city", infoCep.Localidade),
		attribute.String("state", infoCep.Uf),
	)

	weather, err := service.WeatherInfo(ctx, infoCep.Uf, infoCep.Localidade, h.weatherAPIKey)
	if err != nil {
		span.SetAttributes(attribute.String("error", err.Error()))
		span.SetStatus(codes.Error, "Failed to fetch weather info")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "error fetching weather data"})
		return
	}

	response := TemperatureResponse{
		TempC: weather.TempC,
		TempF: weather.TempF,
		TempK: weather.TempK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
