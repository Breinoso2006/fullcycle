package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTemperatureByCep_InvalidFormat(t *testing.T) {
	handler := NewCepHandler("test_api_key")

	testCases := []struct {
		name string
		cep  string
	}{
		{"CEP muito curto", "123"},
		{"CEP muito longo", "123456789"},
		{"CEP com letras", "0100100A"},
		{"CEP vazio", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/weather/"+tc.cep, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			http.HandlerFunc(handler.GetTemperatureByCep).ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusUnprocessableEntity {
				t.Errorf("handler retornou status errado: obteve %v esperado %v", status, http.StatusUnprocessableEntity)
			}

			var response ErrorResponse
			err = json.Unmarshal(rr.Body.Bytes(), &response)
			if err != nil {
				t.Errorf("Erro ao decodificar resposta: %v", err)
			}

			if response.Message != "invalid zipcode" {
				t.Errorf("Mensagem errada: obteve %v esperado 'invalid zipcode'", response.Message)
			}
		})
	}
}

func TestGetTemperatureByCep_NotFound(t *testing.T) {
	handler := NewCepHandler("test_api_key")

	req, err := http.NewRequest("GET", "/weather/99999999", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(handler.GetTemperatureByCep).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler retornou status errado: obteve %v esperado %v", status, http.StatusNotFound)
	}

	var response ErrorResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Erro ao decodificar resposta: %v", err)
	}

	if response.Message != "can not find zipcode" {
		t.Errorf("Mensagem errada: obteve %v esperado 'can not find zipcode'", response.Message)
	}
}

func TestGetTemperatureByCep_MethodNotAllowed(t *testing.T) {
	handler := NewCepHandler("test_api_key")

	req, err := http.NewRequest("POST", "/weather/01001000", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(handler.GetTemperatureByCep).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler retornou status errado: obteve %v esperado %v", status, http.StatusMethodNotAllowed)
	}
}

func TestIsValidCepFormat(t *testing.T) {
	testCases := []struct {
		name     string
		cep      string
		expected bool
	}{
		{"CEP válido sem hífen", "01001000", true},
		{"CEP válido com hífen", "01001-000", true},
		{"CEP inválido curto", "123", false},
		{"CEP inválido longo", "123456789", false},
		{"CEP com letras", "0100100A", false},
		{"CEP vazio", "", false},
		{"CEP com espaços", "  01001000  ", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isValidCepFormat(tc.cep)
			if result != tc.expected {
				t.Errorf("isValidCepFormat(%q) = %v, esperado %v", tc.cep, result, tc.expected)
			}
		})
	}
}
