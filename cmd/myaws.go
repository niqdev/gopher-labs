package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/niqdev/gopher-labs/internal/myaws"
)

func NewMyAwsCmd() *cobra.Command {
	var name string
	command := &cobra.Command{
		Use:   "myaws",
		Short: "Aws examples",
		Run: func(cmd *cobra.Command, args []string) {
			invokeMyAwsCmd(name)
		},
	}
	command.Flags().StringVarP(&name, "name", "n", "", "name of the example")
	return command
}

func init() {
	rootCmd.AddCommand(NewMyAwsCmd())
}

func invokeMyAwsCmd(cmd string) {
	switch cmd {
	case "sqs-write":
		myaws.Send()
	case "sqs-read":
		myaws.Receive()

	default:
		log.Fatalf("invalid command: [%v]", cmd)
	}
}
