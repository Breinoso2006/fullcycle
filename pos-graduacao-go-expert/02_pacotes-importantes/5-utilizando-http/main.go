package main

import (
	"encoding/json"
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
}

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8000", nil) // Multiplexer padrão do Go (nil), global
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {

	// Tratamento de rota
	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Tratamento de parâmetro necessário
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("O parâmetro cep é obrigatório"))
		return
	}

	cepData, err := BuscaCep(cepParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao buscar cep"))
		return
	}

	// Tratamento de Retorno
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cepData)
	// Outra abordagem seria:
	// result, err := json.Marshal(cepData)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte("Erro ao buscar cep"))
	// 	return
	// }
	// w.Write(result)
}

func BuscaCep(cep string) (*ViaCep, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var c ViaCep
	json.Unmarshal(body, &c)

	return &c, nil

}
