package cmd

import (
	"github.com/spf13/cobra"

	"github.com/niqdev/gopher-labs/internal/myspinner"
)

var mySpinnerCmd = &cobra.Command{
	Use:   "myspinner",
	Short: "Spinner examples",
	Run: func(cmd *cobra.Command, args []string) {
		myspinner.Run()
	},
}

func init() {
	rootCmd.AddCommand(mySpinnerCmd)
}
