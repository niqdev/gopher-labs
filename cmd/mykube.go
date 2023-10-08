package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/niqdev/gopher-labs/internal/mykube"
)

func NewMyKubeCmd() *cobra.Command {
	var name string
	command := &cobra.Command{
		Use:   "mykube",
		Short: "Kubernetes examples",
		Run: func(cmd *cobra.Command, args []string) {
			invokeMyKubeCmd(name)
		},
	}
	command.Flags().StringVarP(&name, "name", "n", "", "name of the example")
	return command
}

func init() {
	rootCmd.AddCommand(NewMyKubeCmd())
}

func invokeMyKubeCmd(cmd string) {
	switch cmd {
	case "create":
		mykube.CreateAll()
	case "list":
		mykube.ListPods()
	case "exec":
		mykube.ExecPod()
	case "portforward":
		mykube.PortForward()
	case "job":
		mykube.TailJob()

	default:
		log.Fatalf("invalid command: [%v]", cmd)
	}
}
