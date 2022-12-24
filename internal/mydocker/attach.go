package mydocker

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/dchest/uniuri"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func Attach() {
	imageName := "edgelevel/alpine-xfce-vnc:web-0.6.0"
	containerName := fmt.Sprintf("mydocker-%s", uniuri.NewLen(5))

	ctx := context.Background()

	docker, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("error docker client: %v", err)
	}
	defer docker.Close()

	reader, err := docker.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		log.Fatalf("error image pull: %v", err)
	}
	//defer reader.Close()
	io.Copy(os.Stdout, reader)

	containerConfig := &container.Config{
		Image:        imageName,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		OpenStdin:    true,
		StdinOnce:    true,
		Tty:          true,
		ExposedPorts: nat.PortSet{
			nat.Port("5900/tcp"): {},
			nat.Port("6080/tcp"): {},
		},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"5900/tcp": []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: "5900"}},
			"6080/tcp": []nat.PortBinding{{HostIP: "0.0.0.0", HostPort: "6080"}},
		},
	}

	newContainer, err := docker.ContainerCreate(
		ctx,
		containerConfig,
		hostConfig, // hostConfig
		nil,        // networkingConfig
		nil,        // platform
		containerName)
	if err != nil {
		log.Fatalf("error container create: %v", err)
	}

	containerId := newContainer.ID

	log.Printf("new container: image=%s, name=%s, id=%s", imageName, containerName, containerId)

	if err := docker.ContainerStart(ctx, containerId, types.ContainerStartOptions{}); err != nil {
		log.Fatalf("error container start: %v", err)
	}

	// statusCh, errCh := docker.ContainerWait(ctx, containerId, container.WaitConditionNotRunning)
	// select {
	// case err := <-errCh:
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// case <-statusCh:
	// }

	execCreateResponse, err := docker.ContainerExecCreate(ctx, containerId, types.ExecConfig{
		AttachStdout: true,
		AttachStdin:  true,
		AttachStderr: true,
		Detach:       false,
		Tty:          true,
		Cmd:          []string{"/bin/bash"},
	})
	if err != nil {
		log.Fatalf("error docker exec create: %v", err)
	}

	execAttachResponse, err := docker.ContainerExecAttach(ctx, execCreateResponse.ID, types.ExecStartCheck{
		Tty: true,
	})
	if err != nil {
		log.Fatalf("error docker exec attach: %v", err)
	}
	defer execAttachResponse.Close()

	log.Printf("after attach")

	closeChannel := func() {
		log.Printf("removing docker container: id=%s", containerId)

		if err := docker.ContainerRemove(ctx, containerId, types.ContainerRemoveOptions{Force: true}); err != nil {
			log.Fatalf("error docker remove: %v", err)
		}
	}

	var once sync.Once
	go func() {
		//_, err := stdcopy.StdCopy(os.Stdout, os.Stderr, execAttachResponse.Reader)
		// tty
		_, err := io.Copy(os.Stdout, execAttachResponse.Reader)
		if err != nil {
			log.Fatalf("error copy docker->local: %v", err)
		}
		once.Do(closeChannel)
	}()

	go func() {
		_, err = io.Copy(execAttachResponse.Conn, os.Stdin)
		if err != nil {
			log.Fatalf("error copy local->docker: %v", err)
		}
		once.Do(closeChannel)
	}()
}
