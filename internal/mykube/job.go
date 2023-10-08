package mykube

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/watch"
	applyv1 "k8s.io/client-go/applyconfigurations/core/v1"
)

func TailJob() {
	ctx := context.TODO()
	clientSet := getClientSet()

	// apply namespace
	namespace := namespaceName
	_, err := clientSet.CoreV1().Namespaces().Apply(ctx, applyv1.Namespace(namespace), metav1.ApplyOptions{FieldManager: "application/apply-patch"})
	if err != nil {
		log.Fatalf("error create namespace: %v", err)
	}
	log.Println(fmt.Sprintf("namespace %s successfully applied", namespace))

	// create job
	job, err := clientSet.BatchV1().Jobs(namespace).Create(ctx, buildJob(namespace), metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("error create job: %v", err)
	}
	log.Println(fmt.Sprintf("job %s successfully created", job.Name))

	// watch pod
	labelMap, err := metav1.LabelSelectorAsMap(job.Spec.Selector)
	if err != nil {
		log.Fatalf("error job label: %v", err)
	}
	listOpts := metav1.ListOptions{
		LabelSelector: labels.SelectorFromSet(labelMap).String(),
	}
	// https://stackoverflow.com/questions/56231176/kubernetes-client-go-convert-labelselector-to-label-string
	//listOpts := metav1.ListOptions{LabelSelector: metav1.FormatLabelSelector(job.Spec.Selector)}
	podWatcher, err := clientSet.CoreV1().Pods(namespace).Watch(ctx, listOpts)
	if err != nil {
		log.Fatalf("error watch pod: %v", err)
	}
	for podEvents := range podWatcher.ResultChan() {
		switch podEvents.Type {
		case watch.Modified:
			podEvent, ok := podEvents.Object.(*corev1.Pod)
			if !ok {
				continue
			}
			log.Println(fmt.Sprintf("pod status %v", podEvent.Status.Phase))
			switch podEvent.Status.Phase {
			case corev1.PodRunning, corev1.PodSucceeded:
				podWatcher.Stop()
			}
		default:
			log.Println(fmt.Sprintf("skip pod event: %v", podEvents.Type))
		}
	}

	// get single pod
	pods, err := clientSet.CoreV1().Pods(namespace).List(ctx, listOpts)
	if err != nil {
		log.Fatalf("error list: %v", err)
	}
	podItem := pods.Items[0]
	log.Println(fmt.Sprintf(">>> %s", podItem.Name))

	// stream logs
	stream, err := clientSet.CoreV1().Pods(namespace).
		GetLogs(podItem.Name, &corev1.PodLogOptions{
			//Container: podItem.Spec.Containers[0].Name,
			Follow: true,
		}).
		Stream(ctx)
	if err != nil {
		log.Fatalf("error stream log: %v", err)
	}
	defer stream.Close()

	// print logs
	_, err = io.Copy(os.Stdout, stream)
	if err != nil {
		log.Fatalf("error std log: %v", err)
	}

	// delete job and pod
	backgroundDeletion := metav1.DeletePropagationBackground
	defer clientSet.BatchV1().Jobs(namespace).Delete(ctx, job.Name, metav1.DeleteOptions{
		PropagationPolicy: &backgroundDeletion,
	})
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
			BackoffLimit: int32Ptr(0),
		},
	}
}
