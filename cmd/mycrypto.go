package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/niqdev/gopher-labs/internal/mycrypto"
)

func NewMyCryptoCmd() *cobra.Command {
	var name string
	command := &cobra.Command{
		Use:   "mycrypto",
		Short: "Cryptography examples",
		Run: func(cmd *cobra.Command, args []string) {
			invokeMyCryptoCmd(name)
		},
	}
	command.Flags().StringVarP(&name, "name", "n", "", "name of the example")
	return command
}

func init() {
	rootCmd.AddCommand(NewMyCryptoCmd())
}

func invokeMyCryptoCmd(cmd string) {
	switch cmd {
	case "pgp-message":
		mycrypto.PgpMessage()

	default:
		log.Fatalf("invalid command: [%v]", cmd)
	}
}
