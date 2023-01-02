package cmd

import (
	"github.com/niqdev/gopher-labs/internal/mylog"
	"github.com/spf13/cobra"
)

var myLogCmd = &cobra.Command{
	Use: "mylog",
	Run: func(cmd *cobra.Command, args []string) {
		mylog.ExampleFromDoc()
	},
}

func init() {
	rootCmd.AddCommand(myLogCmd)
}
