package mydocker

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/dchest/uniuri"
	"golang.org/x/crypto/ssh/terminal"

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
	defer reader.Close()

	// suppress output
	// io.Copy(os.Stdout, reader)
	io.Copy(ioutil.Discard, reader)

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

	log.Printf("new container: image=%s | name=%s | id=%s", imageName, containerName, containerId)

	if err := docker.ContainerStart(ctx, containerId, types.ContainerStartOptions{}); err != nil {
		log.Fatalf("error container start: %v", err)
	}

	execCreateResponse, err := docker.ContainerExecCreate(ctx, containerId, types.ExecConfig{
		AttachStdout: true,
		AttachStdin:  true,
		AttachStderr: true,
		Detach:       false,
		Tty:          true,
		Cmd:          []string{"/bin/bash"},
	})
	if err != nil {
		log.Fatalf("error container exec create: %v", err)
	}

	execAttachResponse, err := docker.ContainerExecAttach(ctx, execCreateResponse.ID, types.ExecStartCheck{
		Tty: true,
	})
	if err != nil {
		log.Fatalf("error container exec attach: %v", err)
	}
	defer execAttachResponse.Close()

	closeChannel := func() {
		// (*) log.Printf("removing container container: id=%s", containerId)

		if err := docker.ContainerRemove(ctx, containerId, types.ContainerRemoveOptions{Force: true}); err != nil {
			log.Fatalf("error container remove: %v", err)
		}
	}

	var once sync.Once
	go func() {
		// use with TTY=false only, with TTY=true returns: "Unrecognized input header: 13"
		//_, err := stdcopy.StdCopy(os.Stdout, os.Stderr, execAttachResponse.Reader)

		// TTY
		_, err := io.Copy(os.Stdout, execAttachResponse.Reader)
		if err != nil {
			log.Fatalf("error copy docker->local: %v", err)
		}

		// (*) log.Println("close docker->local")
		once.Do(closeChannel)
	}()

	go func() {
		_, err = io.Copy(execAttachResponse.Conn, os.Stdin)
		if err != nil {
			log.Fatalf("error copy local->docker: %v", err)
		}

		// (*) log.Println("close local->docker")
		once.Do(closeChannel)
	}()

	// https://stackoverflow.com/questions/58732588/accept-user-input-os-stdin-to-container-using-golang-docker-sdk-interactive-co
	// (*) remove commented logging to avoid messing with terminal output
	fd := int(os.Stdin.Fd())
	var oldState *terminal.State
	if terminal.IsTerminal(fd) {
		oldState, err = terminal.MakeRaw(fd)
		if err != nil {
			log.Fatalf("error raw terminal: %v", err)
		}
		defer terminal.Restore(fd, oldState)
	}

	statusCh, errCh := docker.ContainerWait(ctx, containerId, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			log.Fatalf("error container wait: %v", err)
		}
		// (*) log.Println("close container wait errCh")
	case _ = <-statusCh:
		// (*) log.Printf("close container wait statusCh: %v", status.StatusCode)
	}
}
