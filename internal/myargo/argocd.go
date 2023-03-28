package myargo

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/utils/pointer"

	"github.com/argoproj/argo-cd/v2/cmd/argocd/commands/headless"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient"
	argocdclient "github.com/argoproj/argo-cd/v2/pkg/apiclient"
	applicationpkg "github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	argoio "github.com/argoproj/argo-cd/v2/util/io"
	"github.com/argoproj/argo-cd/v2/util/localconfig"
)

// TODO cleanup unused go mod

func NewArgoCdClient(argocdConfigPath string) (apiclient.Client, error) {
	if argocdConfigPath == "" {
		var err error
		argocdConfigPath, err = localconfig.DefaultLocalConfigPath()
		if err != nil {
			return nil, err
		}
	}
	var argocdCliOpts apiclient.ClientOptions
	argocdCliOpts.ConfigPath = argocdConfigPath
	return argocdclient.NewClient(&argocdCliOpts)
}

func ListApplications() {
	var (
		selector     = "app.kubernetes.io/instance"
		appNamespace = "argocd"
	)

	clientOpts := &argocdclient.ClientOptions{}
	argocdClient := headless.NewClientOrDie(clientOpts, &cobra.Command{})
	conn, appIf := argocdClient.NewApplicationClientOrDie()
	defer argoio.Close(conn)
	apps, err := appIf.List(context.Background(), &applicationpkg.ApplicationQuery{
		Selector:     pointer.String(selector),
		AppNamespace: &appNamespace,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(apps.String())
}
