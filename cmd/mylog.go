package cmd

import (
	"github.com/spf13/cobra"

	"github.com/niqdev/gopher-labs/internal/mylog"
)

var myLogCmd = &cobra.Command{
	Use:   "mylog",
	Short: "zap logging examples",
	Run: func(cmd *cobra.Command, args []string) {
		mylog.ExampleFromDoc()
		mylog.ExampleWithColor()
	},
}

func init() {
	rootCmd.AddCommand(myLogCmd)
}
