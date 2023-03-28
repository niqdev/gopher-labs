package cmd

import (
	"github.com/spf13/cobra"

	"github.com/niqdev/gopher-labs/internal/myschema"
)

var mySchemaCmd = &cobra.Command{
	Use:   "myschema",
	Short: "JSON and Yaml schema validation",
	Run: func(cmd *cobra.Command, args []string) {
		myschema.JsonSchemaValidation()
	},
}

func init() {
	rootCmd.AddCommand(mySchemaCmd)
}
