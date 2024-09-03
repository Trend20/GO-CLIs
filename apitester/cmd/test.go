package cmd

import (
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "apitester",
	Short: "A CLI tool that tests APIs and generates API documentation.",
	Long: `A CLI tool that tests APIs and generates API documentation based on the test results.
Support different formats like OpenAPI/Swagger and include automated testing features.`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
}
