package mykube

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/kubectl/pkg/cmd/exec"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

func initPod(namespace string) *corev1.Pod {
	pods := getPods(context.TODO(), namespace, "app.kubernetes.io/name=edgelevel-alpine-xfce-vnc")
	if len(pods) != 1 {
		log.Fatalf("pod alpine-xfce-vnc-* not found or invalid")
	}
	pod := pods[0]
	return &pod
}

// POD_EXAMPLE=$(kubectl -n examples get pods --selector=app.kubernetes.io/name=edgelevel-alpine-xfce-vnc -o name | xargs basename)
// POD_EXAMPLE=$(kubectl -n examples get pods --selector=app.kubernetes.io/name=edgelevel-alpine-xfce-vnc -o name | awk -F '/' '{print $2}')

// kubectl cp ./README.md examples/${POD_EXAMPLE}:/tmp
// kubectl cp ./README.md examples/${POD_EXAMPLE}:/tmp/
// kubectl cp ./README.md examples/${POD_EXAMPLE}:/tmp/readme-copy.md
// kubectl cp ./data examples/${POD_EXAMPLE}:/tmp

func CopyToPod() {
	namespace := namespaceName
	srcPath := "./README.md"
	// overrides destination path if exists
	destPath := "/tmp/readme-copy.md"

	if _, err := os.Stat(srcPath); err != nil {
		log.Fatalf(fmt.Sprintf("error file %s doesn't exist in local filesystem", srcPath))
	}

	restConfig, coreClient := newCoreClient()
	pod := initPod(namespace)

	isTty := false
	executor := exec.DefaultRemoteExecutor{}

	// see checkDestinationIsDir
	requestUrlDir := newRestRequest(coreClient, pod, []string{"test", "-d", destPath}, isTty).URL()
	errDir := executor.Execute(http.MethodPost, requestUrlDir, restConfig, os.Stdin, os.Stdout, os.Stderr, isTty, nil)
	if errDir == nil {
		destPath = destPath + "/" + path.Base(srcPath)
	}

	log.Println(fmt.Sprintf("source path: %s", srcPath))
	log.Println(fmt.Sprintf("destination path: %s/%s:%s", namespace, pod.Name, destPath))

	reader, writer := io.Pipe()

	// create archive
	go func() {
		defer writer.Close()
		err := cpMakeTar(srcPath, destPath, writer)
		cmdutil.CheckErr(err)
	}()

	// extract archive
	var cmdArr []string
	cmdArr = []string{"tar", "-xmf", "-"}
	destDir := path.Dir(destPath)
	if len(destDir) > 0 {
		cmdArr = append(cmdArr, "-C", destDir)
	}

	requestUrlExtract := newRestRequest(coreClient, pod, cmdArr, isTty).URL()
	errExtract := executor.Execute(http.MethodPost, requestUrlExtract, restConfig, reader, os.Stdout, os.Stderr, isTty, nil)
	if errExtract != nil {
		log.Fatalf("error extract all: %v", errExtract)
	}
}

func CopyFromPod() {
	namespace := namespaceName
	srcPath := "/etc/passwd"
	destPath := "./etc-passwd"
	restConfig, coreClient := newCoreClient()
	pod := initPod(namespace)

	reader, writer := io.Pipe()

	isTty := false
	executor := exec.DefaultRemoteExecutor{}

	cmdArr := []string{"tar", "cf", "-", srcPath}
	requestUrlArchive := newRestRequest(coreClient, pod, cmdArr, isTty).URL()

	go func() {
		defer writer.Close()
		errArchive := executor.Execute(http.MethodPost, requestUrlArchive, restConfig, os.Stdin, writer, os.Stderr, isTty, nil)
		cmdutil.CheckErr(errArchive)
	}()

	// see stripLeadingSlash
	prefix := strings.TrimLeft(srcPath, `/\`)
	prefix = path.Clean(prefix)
	prefix = cpStripPathShortcuts(prefix)
	destPath = path.Join(destPath, path.Base(prefix))

	if err := untarAll(reader, destPath, prefix); err != nil {
		log.Fatalf("error untar all: %v", err)
	}
}
