package client

import (
	"context"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

var orchestratorURL = "http://localhost:8081"

func SetOrchestratorURL(url string) {
	orchestratorURL = url
}

func WeatherOrchestratorRequest(ctx context.Context, cep string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, orchestratorURL+"/weather/"+cep, nil)
	if err != nil {
		return nil, err
	}

	// Cliente com instrumentação OTEL
	client := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
