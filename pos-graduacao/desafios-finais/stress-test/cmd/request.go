package cmd

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var request = &cobra.Command{
	Use:   "request",
	Short: "Realiza um teste de carga em um serviço web",
	Long: `Comando para realizar um teste de carga em um serviço web especificado pelo usuário.
	O usuário deve fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetInt("requests")
		concurrency, _ := cmd.Flags().GetInt("concurrency")

		fmt.Printf("Iniciando teste de carga na URL: %s\n", url)
		fmt.Printf("Número total de requests: %d\n", requests)
		fmt.Printf("Número de requests simultâneos: %d\n", concurrency)

		limitChannel := make(chan struct{}, concurrency)
		var wg sync.WaitGroup
		tempo := time.Now()
		dist := make(map[string]int)
		ch := make(chan int)
		ch = <- 

		for i := 0; i < requests; i++ {
			limitChannel <- struct{}{}
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				defer func() { <-limitChannel }()
				status := getUrl(url)
				if dist[status] {
					dist[status] += 1
				} else {
					dist[status] = 1
				}
			}(i)
		}

		
		wg.Wait()
		fmt.Printf("Tempo total de execução: %s\n", time.Since(tempo))
		fmt.Println("All done")
	},
}

func init() {
	rootCmd.AddCommand(request)
	request.Flags().StringP("url", "u", "", "URL do serviço web alvo do teste de carga")
	request.Flags().IntP("requests", "r", 0, "Número total de requests a serem enviados durante o teste de carga")
	request.Flags().IntP("concurrency", "c", 0, "Número de requests simultâneos a serem enviados durante o teste de carga")
}

func getUrl(url string) string {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao fazer a requisição:", err.Error())
		return strconv.Itoa(response.StatusCode)
	}

	return strconv.Itoa(response.StatusCode)
}
