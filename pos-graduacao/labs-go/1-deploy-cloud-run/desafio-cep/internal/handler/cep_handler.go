package handler

import (
	"breinoso2006/desafio-cep/internal/service"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
)

type CepHandler struct {
	weatherAPIKey string
}

type TemperatureResponse struct {
	TempC float64 `json:"temp_C"`
	TempF float64 `json:"temp_F"`
	TempK float64 `json:"temp_K"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewCepHandler(weatherAPIKey string) *CepHandler {
	return &CepHandler{
		weatherAPIKey: weatherAPIKey,
	}
}

func (h *CepHandler) GetTemperatureByCep(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/weather/")
	cep := strings.TrimSpace(path)

	if !isValidCepFormat(cep) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(ErrorResponse{Message: "invalid zipcode"})
		return
	}

	cep = regexp.MustCompile(`\D`).ReplaceAllString(cep, "")

	infoCep, err := service.CepInfo(cep)
	if err != nil {
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

	weather, err := service.WeatherInfo(infoCep.Uf, infoCep.Localidade, h.weatherAPIKey)
	if err != nil {
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

func isValidCepFormat(cep string) bool {
	cep = strings.TrimSpace(cep)
	matched, _ := regexp.MatchString(`^\d{8}$|^\d{5}-?\d{3}$`, cep)
	return matched
}
