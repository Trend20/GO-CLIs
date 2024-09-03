package cmd

import (
	"github.com/spf13/cobra"
)

// command flags
var url string
var body string
var method string
var headers []string

// actual command body
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test an  API endpoint.",
	Long:  `Test a specific API endpoint with specified HTTP method, headers, and body.`,
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	//	assign the flags to the command
	testCmd.Flags().StringVarP(&url, "url", "u", "", "API endpoint URL")
	testCmd.Flags().StringVarP(&method, "method", "m", "GET", "HTTP method")
	testCmd.Flags().StringArrayVarP(&headers, "header", "H", []string{}, "HTTP header")
	testCmd.Flags().StringVarP(&body, "body", "b", "", "HTTP Request body")
}
