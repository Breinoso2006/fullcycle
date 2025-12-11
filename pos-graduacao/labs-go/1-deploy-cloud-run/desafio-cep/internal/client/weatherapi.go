package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Weather struct {
	TempC float64 `json:"temp_c"`
	TempF float64 `json:"temp_f"`
	TempK float64 `json:"temp_k"`
}

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func WeatherApiRequest(uf, localidade, apiKey string) (*Weather, error) {
	query := fmt.Sprintf("%s,%s", localidade, uf)
	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no",
		apiKey, url.QueryEscape(query))

	req, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a requisição: %w", err)
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro na resposta da API: status %d", req.StatusCode)
	}

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o corpo da resposta: %w", err)
	}

	var apiResponse WeatherAPIResponse
	err = json.Unmarshal(res, &apiResponse)
	if err != nil {
		return nil, fmt.Errorf("erro ao converter para struct: %w", err)
	}

	tempC := apiResponse.Current.TempC
	weather := &Weather{
		TempC: tempC,
		TempF: celsiusToFahrenheit(tempC),
		TempK: celsiusToKelvin(tempC),
	}

	return weather, nil
}

func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273
}
