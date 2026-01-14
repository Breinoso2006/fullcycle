package service

import (
	"breinoso2006/labs-go/otel/orchestrator/internal/client"
	"context"
	"fmt"
)

func WeatherInfo(ctx context.Context, uf, localidade, apiKey string) (*client.Weather, error) {
	data, err := client.WeatherApiRequest(ctx, uf, localidade, apiKey)
	if err != nil {
		fmt.Println("Erro:", err)
		return nil, err
	}

	return data, nil
}
