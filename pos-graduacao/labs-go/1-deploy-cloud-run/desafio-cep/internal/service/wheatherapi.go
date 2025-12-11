package service

import (
	"breinoso2006/desafio-cep/internal/client"
	"fmt"
)

func WeatherInfo(uf, localidade, apiKey string) (*client.Weather, error) {
	data, err := client.WeatherApiRequest(uf, localidade, apiKey)
	if err != nil {
		fmt.Println("Erro:", err)
		return nil, err
	}

	return data, nil
}
