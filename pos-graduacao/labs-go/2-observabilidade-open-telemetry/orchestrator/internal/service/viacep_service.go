package service

import (
	"breinoso2006/labs-go/otel/orchestrator/internal/client"
	"context"
	"fmt"
	"regexp"
)

func CepInfo(ctx context.Context, cep string) (*client.ViaCep, error) {
	cep, err := verifyCep(cep)
	if err != nil {
		fmt.Println("Erro:", err)
		return nil, err
	}
	fmt.Println("CEP válido:", cep)

	data, err := client.ViaCepRequest(ctx, cep)
	if err != nil {
		fmt.Println("Erro:", err)
		return nil, err
	}

	return data, nil
}

func verifyCep(cep string) (string, error) {
	cepClean := regexp.MustCompile(`\D`).ReplaceAllString(cep, "")

	if len(cepClean) != 8 {
		return "", fmt.Errorf("CEP inválido: deve conter 8 dígitos")
	}

	if !regexp.MustCompile(`^\d{8}$`).MatchString(cepClean) {
		return "", fmt.Errorf("CEP inválido: deve conter apenas números")
	}

	return cepClean, nil
}

func FormatCep(cep string) string {
	if len(cep) != 8 {
		return cep
	}
	return fmt.Sprintf("%s-%s", cep[:5], cep[5:])
}
