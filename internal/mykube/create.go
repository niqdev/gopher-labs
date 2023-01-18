package mykube

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func CreateAll() {
	ctx := context.TODO()
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("error restConfig: %v", err)
	}

	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		log.Fatalf("error clientSet: %v", err)
	}

	namespace, err := clientSet.CoreV1().Namespaces().Create(ctx, buildNamespace(), metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("error create namespace: %v", err)
	}
	log.Println(fmt.Sprintf("namespace %s successfully created", namespace.Name))

	deployment, err := clientSet.AppsV1().Deployments(namespace.Name).Create(ctx, buildDeployment(), metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("error create deployment: %v", err)
	}
	log.Println(fmt.Sprintf("deployment %s successfully created", deployment.Name))

	service, err := clientSet.CoreV1().Services(namespace.Name).Create(ctx, buildService(), metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("error create service: %v", err)
	}
	log.Println(fmt.Sprintf("service %s successfully created", service.Name))
}

func buildNamespace() *corev1.Namespace {
	return &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "examples",
		},
	}
}

func buildLabels() map[string]string {
	return map[string]string{
		"app.kubernetes.io/name":    "edgelevel-alpine-xfce-vnc",
		"app.kubernetes.io/version": "web-0.6.0",
	}
}

func buildPod() *corev1.Pod {
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "edgelevel-alpine-xfce-vnc",
			Namespace: "examples",
			Labels:    buildLabels(),
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:            "edgelevel-alpine-xfce-vnc",
					Image:           "edgelevel/alpine-xfce-vnc:web-0.6.0",
					ImagePullPolicy: corev1.PullIfNotPresent,
					TTY:             true,
					Stdin:           true,
					Ports: []corev1.ContainerPort{
						{
							Name:          "vnc-svc",
							Protocol:      corev1.ProtocolTCP,
							ContainerPort: 5900,
						},
						{
							Name:          "novnc-svc",
							Protocol:      corev1.ProtocolTCP,
							ContainerPort: 6080,
						},
					},
					Resources: corev1.ResourceRequirements{
						Limits: corev1.ResourceList{
							"memory": resource.MustParse("512Mi"),
						},
						Requests: corev1.ResourceList{
							"cpu":    resource.MustParse("500m"),
							"memory": resource.MustParse("512Mi"),
						},
					},
				},
			},
		},
	}
}

func int32Ptr(i int32) *int32 { return &i }

func buildDeployment() *appsv1.Deployment {

	pod := buildPod()

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "alpine-xfce-vnc-deployment",
			Namespace: "examples",
			Labels:    buildLabels(),
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: buildLabels(),
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: pod.ObjectMeta,
				Spec:       pod.Spec,
			},
		},
	}
}

func buildService() *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "alpine-xfce-vnc-service",
			Namespace: "examples",
			Labels:    buildLabels(),
		},
		Spec: corev1.ServiceSpec{
			Selector: buildLabels(),
			Type:     corev1.ServiceTypeClusterIP,
			Ports: []corev1.ServicePort{
				{
					Name:       "vnc",
					Protocol:   corev1.ProtocolTCP,
					Port:       5900,
					TargetPort: intstr.FromString("vnc-svc"),
				},
				{
					Name:       "novnc",
					Protocol:   corev1.ProtocolTCP,
					Port:       6080,
					TargetPort: intstr.FromString("novnc-svc"),
				},
			},
		},
	}
}
