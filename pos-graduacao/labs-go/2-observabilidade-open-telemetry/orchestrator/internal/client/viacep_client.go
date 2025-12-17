package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
	Erro        bool   `json:"erro,omitempty"`
}

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

var defaultHTTPClient HTTPClient = &http.Client{}

func ViaCepRequest(ctx context.Context, cep string) (*ViaCep, error) {
	tracer := otel.Tracer("service-b")
	ctx, span := tracer.Start(ctx, "viacep-api-call")
	defer span.End()

	span.SetAttributes(attribute.String("cep", cep))

	url := "https://viacep.com.br/ws/" + cep + "/json/"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
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

	var data ViaCep
	err = json.Unmarshal(res, &data)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, fmt.Errorf("erro ao converter para struct: %w", err)
	}

	if data.Erro {
		err := fmt.Errorf("CEP não encontrado")
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}

	span.SetAttributes(
		attribute.String("cidade", data.Localidade),
		attribute.String("uf", data.Uf),
	)

	return &data, nil
}

func ViaCepRequestWithClient(cep string, client HTTPClient) (*ViaCep, error) {
	req, err := client.Get("https://viacep.com.br/ws/" + cep + "/json/")
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

	var data ViaCep
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, fmt.Errorf("erro ao converter para struct: %w", err)
	}

	if data.Erro {
		return nil, fmt.Errorf("CEP não encontrado")
	}

	return &data, nil
}
