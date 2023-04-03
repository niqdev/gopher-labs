package cmd

import (
	"github.com/spf13/cobra"

	"github.com/niqdev/gopher-labs/internal/myargo"
)

func NewMyArgoCmd() *cobra.Command {
	var name string
	command := &cobra.Command{
		Use:   "myargo",
		Short: "Argo examples",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}
	command.Flags().StringVarP(&name, "name", "n", "", "name of the example")

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "list argo-cd applications",
		Run: func(cmd *cobra.Command, args []string) {
			myargo.ListApplications()
		},
	}

	submitCmd := &cobra.Command{
		Use:   "submit",
		Short: "submit argo workflow",
		Run: func(cmd *cobra.Command, args []string) {
			myargo.SubmitWorkflow()
		},
	}

	command.AddCommand(listCmd)
	command.AddCommand(submitCmd)
	return command
}

func init() {
	rootCmd.AddCommand(NewMyArgoCmd())
}
