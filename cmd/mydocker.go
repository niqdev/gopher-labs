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
			invokeCmd(name)
		},
	}
	command.Flags().StringVarP(&name, "name", "n", "", "name of the example")
	return command
}

func init() {
	rootCmd.AddCommand(NewMydockerCmd())
}

func invokeCmd(cmd string) {
	switch cmd {
	case "e1":
		mydocker.Example1()

	default:
		log.Fatalf("invalid command: [%v]", cmd)
	}
}
