package cmd

import (
	"github.com/niqdev/gopher-labs/internal/myspinner"
	"github.com/spf13/cobra"
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
