package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/niqdev/gopher-labs/internal/myhttp"
)

func NewMyHttpCmd() *cobra.Command {
	var name string
	command := &cobra.Command{
		Use:   "myhttp",
		Short: "HTTP examples",
		Run: func(cmd *cobra.Command, args []string) {
			invokeMyHttpCmd(name)
		},
	}
	command.Flags().StringVarP(&name, "name", "n", "", "name of the example")
	return command
}

func init() {
	rootCmd.AddCommand(NewMyHttpCmd())
}

func invokeMyHttpCmd(cmd string) {
	switch cmd {
	case "client":
		myhttp.SimpleHttpRequest()
		myhttp.RetryHttpRequest()
	case "server":
		myhttp.StartServer()
	case "ws-server":
		myhttp.WebsocketServer()
	case "ws-client":
		myhttp.WebsocketClient()

	default:
		log.Fatalf("invalid command: [%v]", cmd)
	}
}
