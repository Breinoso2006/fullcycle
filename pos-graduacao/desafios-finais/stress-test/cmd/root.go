package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stress-test",
	Short: "CLI para realizar testes de carga em um serviço web",
	Long: `Sistema CLI em Go para realizar testes de carga em um serviço web. 
	O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
