package mykube

import (
	"context"
	"fmt"
	"log"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TailJob() {
	ctx := context.TODO()
	clientSet := getClientSet()

	namespace, err := clientSet.CoreV1().Namespaces().Create(ctx, buildNamespace(), metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("error create namespace: %v", err)
	}
	log.Println(fmt.Sprintf("namespace %s successfully created", namespace.Name))

	job, err := clientSet.BatchV1().Jobs(namespace.Name).Create(ctx, buildJob(namespace.Name), metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("error create job: %v", err)
	}
	log.Println(fmt.Sprintf("job %s successfully created", job.Name))
}

func buildJob(namespace string) *batchv1.Job {
	name := "whalesay-hello"

	return &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            name,
							Image:           "docker/whalesay",
							ImagePullPolicy: corev1.PullIfNotPresent,
							Command:         []string{"cowsay", "hello world"},
						},
					},
					RestartPolicy: corev1.RestartPolicyNever,
				},
			},
			BackoffLimit: int32Ptr(1),
		},
	}
}
