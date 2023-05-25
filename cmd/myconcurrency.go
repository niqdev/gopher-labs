package cmd

import (
	"github.com/spf13/cobra"

	"github.com/niqdev/gopher-labs/internal/myconcurrency"
)

func NewMyConcurrencyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "myconcurrency",
		Short: "Concurrency examples",
		Run: func(cmd *cobra.Command, args []string) {
			myconcurrency.ColorFilter()
			// TODO
			//myconcurrency.ColorFilterGroup()
		},
	}
}

func init() {
	rootCmd.AddCommand(NewMyConcurrencyCmd())
}
