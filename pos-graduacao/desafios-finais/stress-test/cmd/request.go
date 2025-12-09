package cmd

import (
	"fmt"
	"net/http"
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
		var mu sync.Mutex
		startTime := time.Now()
		statusDist := make(map[int]int)
		totalRequests := 0

		for i := 0; i < requests; i++ {
			limitChannel <- struct{}{}
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				defer func() { <-limitChannel }()
				
				statusCode := makeRequest(url)
				
				mu.Lock()
				statusDist[statusCode]++
				totalRequests++
				mu.Unlock()
			}(i)
		}

		wg.Wait()
		totalTime := time.Since(startTime)
		
		generateReport(totalTime, totalRequests, statusDist)
	},
}

func init() {
	rootCmd.AddCommand(request)
	request.Flags().StringP("url", "u", "", "URL do serviço web alvo do teste de carga")
	request.Flags().IntP("requests", "r", 0, "Número total de requests a serem enviados durante o teste de carga")
	request.Flags().IntP("concurrency", "c", 0, "Número de requests simultâneos a serem enviados durante o teste de carga")
}

func generateReport(totalTime time.Duration, totalRequests int, statusDist map[int]int) {
	fmt.Println("\n========== RELATÓRIO DO TESTE DE CARGA ==========")
	fmt.Printf("Tempo total gasto na execução: %s\n", totalTime)
	fmt.Printf("Quantidade total de requests realizados: %d\n", totalRequests)
		
	fmt.Println("\nDistribuição de códigos de status HTTP:")
	for status, count := range statusDist {
		if status == 0 {
			fmt.Printf("  Erro de conexão/timeout: %d requests\n", count)
		} else {
			fmt.Printf("  Status %d: %d requests\n", status, count)
		}
	}
	fmt.Println("=================================================")
}

func makeRequest(url string) int {
	response, err := http.Get(url)
	if err != nil {
		// Retorna 0 para erros de conexão/timeout
		return 0
	}
	defer response.Body.Close()
	
	return response.StatusCode
}
