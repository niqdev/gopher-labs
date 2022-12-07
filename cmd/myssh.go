package cmd

import (
	"log"
	"os"

	"github.com/niqdev/gopher-labs/internal/myssh"
	"github.com/spf13/cobra"
)

func NewMysshCmd() *cobra.Command {
	command := &cobra.Command{
		Use:   "myssh",
		Short: "Simple SSH server and client",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.HelpFunc()(cmd, args)
				os.Exit(1)
			}
		},
	}

	serverCmd := &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println(myssh.RunServer())
		},
	}

	clientCmd := &cobra.Command{
		Use: "client",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println(myssh.RunClient())
		},
	}

	command.AddCommand(serverCmd)
	command.AddCommand(clientCmd)
	return command
}

func init() {
	rootCmd.AddCommand(NewMysshCmd())
}
