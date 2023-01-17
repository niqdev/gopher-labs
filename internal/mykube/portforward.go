package mykube

import (
	"context"
	"log"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func PortForward() {
	ctx := context.Background()
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("error kubeconfig: %v", err)
	}

	_, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("error clientset: %v", err)
	}

	pods := getPods(ctx, "examples", "app.kubernetes.io/name=box-edgelevel-alpine-xfce-vnc")
	if len(pods) != 1 {
		log.Fatalf("pod alpine-xfce-vnc-* not found or invalid")
	}
	podName := pods[0].ObjectMeta.Name
	log.Println(podName)
}
