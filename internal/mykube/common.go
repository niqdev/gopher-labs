package mykube

import (
	"fmt"
	"k8s.io/client-go/rest"
	"log"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func newClientSet() *kubernetes.Clientset {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("error restConfig: %v", err)
	}

	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		log.Fatalf("error clientSet: %v", err)
	}
	return clientSet
}

func newCoreClient() (*rest.Config, *corev1client.CoreV1Client) {
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		clientcmd.NewDefaultClientConfigLoadingRules(),
		&clientcmd.ConfigOverrides{},
	)

	// not used
	namespace, _, err := kubeconfig.Namespace()
	if err != nil {
		log.Fatalf("error namespace: %v", err)
	}
	log.Println(fmt.Sprintf("current namespace: %s", namespace))

	restConfig, err := kubeconfig.ClientConfig()
	if err != nil {
		log.Fatalf("error restConfig: %v", err)
	}

	coreClient, err := corev1client.NewForConfig(restConfig)
	if err != nil {
		log.Fatalf("error coreClient: %v", err)
	}
	return restConfig, coreClient
}
