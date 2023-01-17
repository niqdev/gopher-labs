package mykube

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/portforward"
	"k8s.io/client-go/transport/spdy"
	"k8s.io/client-go/util/homedir"
)

func PortForward() {
	var wg sync.WaitGroup
	wg.Add(1)

	ctx := context.Background()
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("error restConfig: %v", err)
	}

	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		log.Fatalf("error clientSet: %v", err)
	}

	pods := getPods(ctx, "examples", "app.kubernetes.io/name=box-edgelevel-alpine-xfce-vnc")
	if len(pods) != 1 {
		log.Fatalf("pod alpine-xfce-vnc-* not found or invalid")
	}
	pod := pods[0]
	log.Println(pod.Name)

	url := clientSet.CoreV1().RESTClient().Post().
		Resource("pods").
		Namespace(pod.Namespace).
		Name(pod.Name).
		SubResource("portforward").URL()

	transport, upgrader, err := spdy.RoundTripperFor(restConfig)
	if err != nil {
		log.Fatalf("error round tripper: %v", err)
	}
	dialer := spdy.NewDialer(upgrader, &http.Client{Transport: transport}, http.MethodPost, url)

	var portBindings []string
	ports := []string{"5900", "6080"}
	for _, port := range ports {
		if err := verifyOpenPort(port); err == nil {
			portBindings = append(portBindings, fmt.Sprintf("%s:%s", port, port))
		} else {
			// warning unable to bind
			log.Println(err)
		}
	}

	stopChan := ctx.Done()
	readyChan := make(chan struct{}, 1)
	streams := genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr,
	}

	forwarder, err := portforward.New(dialer, portBindings, stopChan, readyChan, streams.Out, streams.ErrOut)
	if err != nil {
		log.Fatalf("error portforward: %v", err)
	}

	go func() {
		err = forwarder.ForwardPorts()
		if err != nil {
			log.Fatal(err)
		}
	}()
	for range readyChan {
	}

	wg.Wait()
}

func verifyOpenPort(port string) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("[::]:%s", port))
	if err != nil {
		return fmt.Errorf("unable to listen on port %s: %v", port, err)
	}

	if err := listener.Close(); err != nil {
		return fmt.Errorf("failed to close port %s: %v", port, err)
	}

	return nil
}
