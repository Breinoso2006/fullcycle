package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func ViaCepRequest(cep string) (*ViaCep, error) {
	return ViaCepRequestWithClient(cep, defaultHTTPClient)
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
