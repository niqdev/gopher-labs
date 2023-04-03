package myargo

import (
	"context"
	"fmt"
	"k8s.io/client-go/util/homedir"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/utils/pointer"

	wfv1 "github.com/argoproj/argo-workflows/v3/pkg/apis/workflow/v1alpha1"
	wfclientset "github.com/argoproj/argo-workflows/v3/pkg/client/clientset/versioned"
)

// https://github.com/argoproj/argo-workflows/blob/master/examples/example-golang/main.go

var helloWorldWorkflow = wfv1.Workflow{
	ObjectMeta: metav1.ObjectMeta{
		GenerateName: "hello-world-",
	},
	Spec: wfv1.WorkflowSpec{
		Entrypoint: "whalesay",
		Templates: []wfv1.Template{
			{
				Name: "whalesay",
				Container: &corev1.Container{
					Image:   "docker/whalesay:latest",
					Command: []string{"cowsay", "hello world"},
				},
			},
		},
	},
}

func SubmitWorkflow() {
	ctx := context.Background()

	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(fmt.Errorf("error restConfig: %v", err))
	}

	namespace := "examples"
	workflowClient := wfclientset.NewForConfigOrDie(restConfig).ArgoprojV1alpha1().Workflows(namespace)

	// submit the hello world workflow
	createdWf, err := workflowClient.Create(ctx, &helloWorldWorkflow, metav1.CreateOptions{})
	if err != nil {
		panic(fmt.Errorf("error create workflow: %v", err))
	}
	fmt.Printf("Workflow %s submitted\n", createdWf.Name)

	// wait for the workflow to complete
	fieldSelector := fields.ParseSelectorOrDie(fmt.Sprintf("metadata.name=%s", createdWf.Name))
	watchIf, err := workflowClient.Watch(ctx, metav1.ListOptions{FieldSelector: fieldSelector.String(), TimeoutSeconds: pointer.Int64(180)})
	defer watchIf.Stop()
	for next := range watchIf.ResultChan() {
		wf, ok := next.Object.(*wfv1.Workflow)
		if !ok {
			continue
		}
		if !wf.Status.FinishedAt.IsZero() {
			fmt.Printf("Workflow %s %s at %v. Message: %s.\n", wf.Name, wf.Status.Phase, wf.Status.FinishedAt, wf.Status.Message)
			break
		}
	}
}
