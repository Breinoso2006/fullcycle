package client

import (
	"testing"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	testCases := []struct {
		name     string
		celsius  float64
		expected float64
	}{
		{"Zero Celsius", 0, 32},
		{"Temperatura ambiente", 25, 77},
		{"Água fervendo", 100, 212},
		{"Temperatura negativa", -10, 14},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := celsiusToFahrenheit(tc.celsius)
			if result != tc.expected {
				t.Errorf("celsiusToFahrenheit(%v) = %v, esperado %v", tc.celsius, result, tc.expected)
			}
		})
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	testCases := []struct {
		name     string
		celsius  float64
		expected float64
	}{
		{"Zero Celsius", 0, 273},
		{"Temperatura ambiente", 25, 298},
		{"Água fervendo", 100, 373},
		{"Zero absoluto", -273, 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := celsiusToKelvin(tc.celsius)
			if result != tc.expected {
				t.Errorf("celsiusToKelvin(%v) = %v, esperado %v", tc.celsius, result, tc.expected)
			}
		})
	}
}
