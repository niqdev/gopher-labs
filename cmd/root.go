package cmd

import (
	"log"

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
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
