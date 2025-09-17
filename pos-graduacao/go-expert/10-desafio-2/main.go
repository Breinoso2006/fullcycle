package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ViaCepResponse struct {
	CEP          string `json:"cep"`
	State        string `json:"uf"`
	City         string `json:"localidade"`
	Neighborhood string `json:"bairro"`
	Street       string `json:"logradouro"`
}

type BrasilApiResponse struct {
	CEP          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

func main() {
	cep := "12030-160"
	channelVia := make(chan ViaCepResponse)
	channelBrasil := make(chan BrasilApiResponse)

	go RequestViaCep(cep, channelVia)
	go RequestBrasilApi(cep, channelBrasil)

	select {
	case value := <-channelVia:
		fmt.Printf("ViaCepApi\n %+v\n", value)
	case value := <-channelBrasil:
		fmt.Printf("BrasilApi\n %+v\n", value)
	case <-time.After(time.Second):
		println("[Error] Timeout")
	}
}

func RequestViaCep(cep string, data chan ViaCepResponse) error {
	// time.Sleep(time.Second * 2)
	req, err := http.NewRequest("GET", "http://viacep.com.br/ws/"+cep+"/json/", nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return err
	}

	var address ViaCepResponse
	err = json.Unmarshal(body, &address)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return err
	}
	data <- address

	return nil
}

func RequestBrasilApi(cep string, data chan BrasilApiResponse) error {
	// time.Sleep(time.Second * 2)
	req, err := http.NewRequest("GET", "https://brasilapi.com.br/api/cep/v1/"+cep, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return err
	}

	var address BrasilApiResponse
	err = json.Unmarshal(body, &address)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return err
	}
	data <- address

	return nil
}
