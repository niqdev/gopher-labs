package cmd

import (
	"github.com/spf13/cobra"

	"github.com/niqdev/gopher-labs/internal/mydb"
)

var myDbCmd = &cobra.Command{
	Use:   "mydb",
	Short: "SQLite CRUD examples",
	Run: func(cmd *cobra.Command, args []string) {
		mydb.SQLiteCrud()
	},
}

func init() {
	rootCmd.AddCommand(myDbCmd)
}
