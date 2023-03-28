package cmd

import (
	"github.com/spf13/cobra"

	"github.com/niqdev/gopher-labs/internal/myconfig"
)

var myConfigCmd = &cobra.Command{
	Use:   "myconfig",
	Short: "Load configs from file",
	Run: func(cmd *cobra.Command, args []string) {
		myconfig.Load()
		myconfig.Format()
	},
}

func init() {
	rootCmd.AddCommand(myConfigCmd)
}
