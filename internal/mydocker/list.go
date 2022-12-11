package mydocker

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// https://github.com/nanobox-io/golang-docker-client
func List() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	log.Println("IMAGES:")
	for _, image := range images {
		fmt.Println(image.ID)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	log.Println("CONTAINERS:")
	for _, container := range containers {
		fmt.Println(container.ID)
	}
}
