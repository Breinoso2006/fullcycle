package client

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockHTTPClient struct {
	response *http.Response
	err      error
}

func (m *mockHTTPClient) Get(url string) (*http.Response, error) {
	return m.response, m.err
}

func TestViaCepRequestWithClient_Success(t *testing.T) {
	mockResponse := ViaCep{
		Cep:        "01001-000",
		Logradouro: "Praça da Sé",
		Bairro:     "Sé",
		Localidade: "São Paulo",
		Uf:         "SP",
		Estado:     "São Paulo",
		Regiao:     "Sudeste",
		Ibge:       "3550308",
		Ddd:        "11",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	client := &http.Client{}
	resp, _ := client.Get(server.URL)
	mockClient := &mockHTTPClient{response: resp}

	result, err := ViaCepRequestWithClient("01001000", mockClient)

	if err != nil {
		t.Errorf("ViaCepRequestWithClient retornou erro inesperado: %v", err)
	}
	if result == nil {
		t.Fatal("ViaCepRequestWithClient retornou resultado nil")
	}
	if result.Localidade != "São Paulo" {
		t.Errorf("Esperado Localidade = 'São Paulo', obtido '%s'", result.Localidade)
	}
}

func TestViaCepRequestWithClient_CepNotFound(t *testing.T) {
	mockResponse := ViaCep{
		Erro: true,
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	client := &http.Client{}
	resp, _ := client.Get(server.URL)
	mockClient := &mockHTTPClient{response: resp}

	_, err := ViaCepRequestWithClient("99999999", mockClient)

	if err == nil {
		t.Error("ViaCepRequestWithClient deveria retornar erro para CEP não encontrado")
	}
	if err != nil && !strings.Contains(err.Error(), "CEP não encontrado") {
		t.Errorf("Erro esperado 'CEP não encontrado', obtido: %v", err)
	}
}

func TestViaCepRequestWithClient_HTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := &http.Client{}
	resp, _ := client.Get(server.URL)
	mockClient := &mockHTTPClient{response: resp}

	_, err := ViaCepRequestWithClient("01001000", mockClient)

	if err == nil {
		t.Error("ViaCepRequestWithClient deveria retornar erro para status HTTP != 200")
	}
}

func TestViaCepRequestWithClient_InvalidJSON(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader("invalid json")),
	}
	mockClient := &mockHTTPClient{response: resp}

	_, err := ViaCepRequestWithClient("01001000", mockClient)

	if err == nil {
		t.Error("ViaCepRequestWithClient deveria retornar erro para JSON inválido")
	}
}

func TestViaCepRequest_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Pulando teste de integração")
	}

	testCases := []struct {
		name        string
		cep         string
		shouldError bool
	}{
		{
			name:        "CEP válido",
			cep:         "01001000",
			shouldError: false,
		},
		{
			name:        "CEP inválido",
			cep:         "99999999",
			shouldError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ViaCepRequest(tc.cep)

			if tc.shouldError {
				if err == nil {
					t.Errorf("ViaCepRequest(%q) deveria retornar erro mas não retornou", tc.cep)
				}
			} else {
				if err != nil {
					t.Errorf("ViaCepRequest(%q) retornou erro inesperado: %v", tc.cep, err)
				}
				if result == nil {
					t.Errorf("ViaCepRequest(%q) retornou resultado nil", tc.cep)
				}
				if result != nil && result.Localidade == "" {
					t.Errorf("ViaCepRequest(%q) retornou localidade vazia", tc.cep)
				}
			}
		})
	}
}
