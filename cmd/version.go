package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	v "github.com/niqdev/gopher-labs/internal"
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
