package cmd

import (
	"github.com/niqdev/gopher-labs/internal/myconfig"
	"github.com/spf13/cobra"
)

var myConfigCmd = &cobra.Command{
	Use:   "myconfig",
	Short: "load configs from file",
	Run: func(cmd *cobra.Command, args []string) {
		myconfig.Load()
	},
}

func init() {
	rootCmd.AddCommand(myConfigCmd)
}