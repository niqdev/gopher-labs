package myargo

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"k8s.io/utils/pointer"

	"github.com/argoproj/argo-cd/v2/cmd/argocd/commands/headless"
	argocdclient "github.com/argoproj/argo-cd/v2/pkg/apiclient"
	applicationpkg "github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	argoio "github.com/argoproj/argo-cd/v2/util/io"
)

func ListApplications() {
	// without env portforward uses by default "~/.kube/config"
	//os.Setenv(clientcmd.RecommendedConfigPathEnvVar, "/my/kubeconfig/path")

	// out-of-cluster kube config
	clientOpts := &argocdclient.ClientOptions{Core: true}
	c := &cobra.Command{}
	c.SetContext(context.Background())

	// starts local server and forward requests to cluster
	argoCdClient := headless.NewClientOrDie(clientOpts, c)
	conn, appIf := argoCdClient.NewApplicationClientOrDie()
	defer argoio.Close(conn)

	var (
		selector     = "" // all applications
		appNamespace = "argocd"
	)
	apps, err := appIf.List(context.Background(), &applicationpkg.ApplicationQuery{
		Selector:     pointer.String(selector),
		AppNamespace: &appNamespace,
	})
	if err != nil {
		panic(err)
	}
	for _, app := range apps.Items {
		fmt.Println(app.QualifiedName())
	}
}
