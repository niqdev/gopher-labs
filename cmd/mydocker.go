package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/niqdev/gopher-labs/internal/mydocker"
)

func NewMydockerCmd() *cobra.Command {
	var name string
	command := &cobra.Command{
		Use:   "mydocker",
		Short: "Docker examples",
		Run: func(cmd *cobra.Command, args []string) {
			log.Printf("NAME %s", name)
			mydocker.Example1()
		},
	}
	command.Flags().StringVarP(&name, "name", "n", "", "name of the example")
	// name := command.Flags().StringP("name", "n", "", "name of the example")
	return command
}

func init() {
	rootCmd.AddCommand(NewMydockerCmd())
}
