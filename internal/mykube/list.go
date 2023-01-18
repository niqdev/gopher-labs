package mykube

import (
	"context"
	"fmt"
	"log"
	"os"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func ListPods() {
	// all namespaces and labels
	pods := getPods(context.TODO(), corev1.NamespaceAll, "")

	for _, pod := range pods {
		log.Println(fmt.Sprintf("%s | %s", pod.ObjectMeta.Namespace, pod.ObjectMeta.Name))
		for keyLabel, valueLabel := range pod.ObjectMeta.Labels {
			log.Println(fmt.Sprintf("  %s = %s", keyLabel, valueLabel))
		}
	}
}

func getPods(ctx context.Context, namespace string, podSelector string) []corev1.Pod {
	kubeconfig := os.Getenv("HOME") + "/.kube/config"

	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("error kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		log.Fatalf("error clientset: %v", err)
	}

	pods, err := clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{LabelSelector: podSelector})
	if err != nil {
		log.Fatalf("error list: %v", err)
	}

	return pods.Items
}
