package mykube

import (
	"context"
	"log"
	"os"

	"golang.org/x/crypto/ssh/terminal"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	corev1client "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

func ExecShellPod() {
	namespace := corev1.NamespaceDefault
	ctx := context.TODO()
	restConfig, coreClient := newCoreClient()

	// creates a busybox Pod: by running "cat", the Pod will sit and do nothing
	pod, err := coreClient.Pods(namespace).Create(ctx, &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "busybox",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"cat"},
					Stdin:   true,
				},
			},
		},
	}, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("error pod creation: %v", err)
	}

	// deletes the Pod on exit
	defer coreClient.Pods(namespace).Delete(ctx, pod.Name, metav1.DeleteOptions{})

	// waits for the Pod to indicate Ready == True
	watcher, err := coreClient.Pods(namespace).Watch(ctx, metav1.SingleObject(pod.ObjectMeta))
	if err != nil {
		log.Fatalf("error pod watcher: %v", err)
	}

	for event := range watcher.ResultChan() {
		switch event.Type {
		case watch.Modified:
			pod = event.Object.(*corev1.Pod)

			// if the Pod contains a status condition Ready == True, stop watching
			for _, cond := range pod.Status.Conditions {
				if cond.Type == corev1.PodReady &&
					cond.Status == corev1.ConditionTrue {
					watcher.Stop()
				}
			}

		default:
			log.Fatalf("error event type: %s", event.Type)
		}
	}

	// exec remote shell
	restRequest := newRestRequest(coreClient, pod, []string{"/bin/sh"}, true)

	exec, err := remotecommand.NewSPDYExecutor(restConfig, "POST", restRequest.URL())
	if err != nil {
		log.Fatalf("error spdy executor: %v", err)
	}

	// put the terminal into raw mode to prevent it echoing characters twice
	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		log.Fatalf("error raw terminal: %v", err)
	}
	defer terminal.Restore(0, oldState)

	// connect this process std{in,out,err} to the remote shell process
	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Tty:    true,
	})
	if err != nil {
		log.Fatalf("error exec stream: %v", err)
	}
}

func newRestRequest(coreClient *corev1client.CoreV1Client, pod *corev1.Pod, commands []string, isTty bool) *rest.Request {
	return coreClient.RESTClient().
		Post().
		Namespace(pod.Namespace).
		Resource("pods").
		Name(pod.Name).
		SubResource("exec").
		VersionedParams(&corev1.PodExecOptions{
			Container: pod.Spec.Containers[0].Name,
			Command:   commands,
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       isTty,
		}, scheme.ParameterCodec)
}
