package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/niqdev/gopher-labs/internal/mydocker"
)

func NewMyDockerCmd() *cobra.Command {
	var name string
	command := &cobra.Command{
		Use:   "mydocker",
		Short: "Docker examples",
		Run: func(cmd *cobra.Command, args []string) {
			invokeMyDockerCmd(name)
		},
	}
	command.Flags().StringVarP(&name, "name", "n", "", "name of the example")
	return command
}

func init() {
	rootCmd.AddCommand(NewMyDockerCmd())
}

func invokeMyDockerCmd(cmd string) {
	switch cmd {
	case "run":
		mydocker.Run()
	case "list":
		mydocker.List()
	case "attach":
		mydocker.Attach()

	default:
		log.Fatalf("invalid command: [%v]", cmd)
	}
}
