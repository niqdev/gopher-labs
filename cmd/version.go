package cmd

import (
	"fmt"

	v "github.com/niqdev/gopher-labs/internal"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(v.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
