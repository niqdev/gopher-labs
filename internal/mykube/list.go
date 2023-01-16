package mykube

import (
	"context"
	"fmt"
	"log"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func ListPods() {
	kubeconfig := os.Getenv("HOME") + "/.kube/config"

	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("error kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		log.Fatalf("error clientset: %v", err)
	}

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, pod := range pods.Items {
		log.Println(fmt.Sprintf("%s | %s", pod.ObjectMeta.Namespace, pod.ObjectMeta.Name))
	}
}
