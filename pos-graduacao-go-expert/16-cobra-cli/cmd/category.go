/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		// name, _ := cmd.Flags().GetString("name")
		// 	fmt.Println("category called with name:", category)
		// 	active, _ := cmd.Flags().GetBool("active")
		// 	fmt.Println("category called with active:", active)
		// 	id, _ := cmd.Flags().GetInt16("id")
		// 	fmt.Println("category called with id:", id)
		// },
		// PreRun: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("PreRun: Command execution started.")
		// },
		// PostRun: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("PostRun: Command execution completed.")
		// },
		// RunE: func(cmd *cobra.Command, args []string) error {
		// 	if category == "" {
		// 		return fmt.Errorf("RunE: An error occurred while executing the command")
		// 	}
		// 	return nil
	},
}

var category string

func init() {
	rootCmd.AddCommand(categoryCmd)
	// categoryCmd.PersistentFlags().StringP("name", "n", "", "Name of the category")
	// categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "", "Name of the category")
	// categoryCmd.PersistentFlags().BoolP("active", "a", false, "Set category as active")
	// categoryCmd.PersistentFlags().Int16P("id", "i", 0, "ID of the category")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// categoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// categoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
