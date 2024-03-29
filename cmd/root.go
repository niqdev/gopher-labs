package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "labs",
	Short: "labs short description",
	Long: `
  .__        ___.           
  |  | _____ \_ |__   ______
  |  | \__  \ | __ \ /  ___/
  |  |__/ __ \| \_\ \\___ \ 
  |____(____  /___  /____  >
            \/    \/     \/`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	// removes timestamps
	log.SetFlags(0)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
