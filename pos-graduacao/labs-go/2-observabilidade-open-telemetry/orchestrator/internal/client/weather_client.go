package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
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

func WeatherApiRequest(ctx context.Context, uf, localidade, apiKey string) (*Weather, error) {
	tracer := otel.Tracer("service-b")
	ctx, span := tracer.Start(ctx, "weatherapi-api-call")
	defer span.End()

	span.SetAttributes(
		attribute.String("cidade", localidade),
		attribute.String("uf", uf),
	)

	query := fmt.Sprintf("%s,%s", localidade, uf)
	weatherURL := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no",
		apiKey, url.QueryEscape(query))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, weatherURL, nil)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}

	client := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	resp, err := client.Do(req)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, fmt.Errorf("erro ao fazer a requisição: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("erro na resposta da API: status %d", resp.StatusCode)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, fmt.Errorf("erro ao ler o corpo da resposta: %w", err)
	}

	var apiResponse WeatherAPIResponse
	err = json.Unmarshal(res, &apiResponse)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, fmt.Errorf("erro ao converter para struct: %w", err)
	}

	tempC := apiResponse.Current.TempC
	weather := &Weather{
		TempC: tempC,
		TempF: celsiusToFahrenheit(tempC),
		TempK: celsiusToKelvin(tempC),
	}

	span.SetAttributes(
		attribute.Float64("temp_celsius", tempC),
		attribute.Float64("temp_fahrenheit", weather.TempF),
		attribute.Float64("temp_kelvin", weather.TempK),
	)

	return weather, nil
}

func celsiusToFahrenheit(celsius float64) float64 {
	return celsius*1.8 + 32
}

func celsiusToKelvin(celsius float64) float64 {
	return celsius + 273
}
