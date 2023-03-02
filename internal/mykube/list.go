package mykube

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
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

	listPodsForService(context.TODO(), corev1.NamespaceAll)
}

func getClientSet() *kubernetes.Clientset {
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

func getPods(ctx context.Context, namespace string, podSelector string) []corev1.Pod {
	clientSet := getClientSet()

	pods, err := clientSet.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{LabelSelector: podSelector})
	if err != nil {
		log.Fatalf("error list: %v", err)
	}

	return pods.Items
}

func listPodsForService(ctx context.Context, namespace string) {
	clientSet := getClientSet()

	services, err := clientSet.CoreV1().Services(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		log.Fatalf("error list service: %v", err)
	}

	for _, service := range services.Items {
		log.Println(fmt.Sprintf("pods for service: name=%s, labels=%v", service.Name, service.GetLabels()))

		labelSet := labels.Set(service.Spec.Selector)
		listOptions := metav1.ListOptions{LabelSelector: labelSet.AsSelector().String()}

		if pods, err := clientSet.CoreV1().Pods(namespace).List(ctx, listOptions); err != nil {
			log.Fatalf("error list pods: %v", err)
		} else {
			for _, pod := range pods.Items {
				log.Println(fmt.Sprintf("* %s", pod.GetName()))
			}
		}
	}
}
