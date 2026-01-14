package service

import (
	"testing"
)

func TestVerifyCep_Valid(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"CEP válido sem hífen", "01001000", "01001000"},
		{"CEP válido com hífen", "01001-000", "01001000"},
		{"CEP com espaços", "  01001-000  ", "01001000"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := verifyCep(tc.input)
			if err != nil {
				t.Errorf("verifyCep(%q) retornou erro inesperado: %v", tc.input, err)
			}
			if result != tc.expected {
				t.Errorf("verifyCep(%q) = %q, esperado %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestVerifyCep_Invalid(t *testing.T) {
	testCases := []struct {
		name  string
		input string
	}{
		{"CEP muito curto", "123"},
		{"CEP muito longo", "123456789"},
		{"CEP com letras", "0100100A"},
		{"CEP vazio", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := verifyCep(tc.input)
			if err == nil {
				t.Errorf("verifyCep(%q) deveria retornar erro mas não retornou", tc.input)
			}
		})
	}
}

func TestFormatCep(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"CEP válido", "01001000", "01001-000"},
		{"CEP inválido curto", "123", "123"},
		{"CEP inválido longo", "123456789", "123456789"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := FormatCep(tc.input)
			if result != tc.expected {
				t.Errorf("FormatCep(%q) = %q, esperado %q", tc.input, result, tc.expected)
			}
		})
	}
}
