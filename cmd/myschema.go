package cmd

import (
	"github.com/niqdev/gopher-labs/internal/myschema"
	"github.com/spf13/cobra"
)

var mySchemaCmd = &cobra.Command{
	Use: "myschema",
	Run: func(cmd *cobra.Command, args []string) {
		myschema.JsonSchemaValidation()
	},
}

func init() {
	rootCmd.AddCommand(mySchemaCmd)
}
