package local

import (
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

func RemoveContainer(container_id string) error {
	ctx := context.Background()
	cli := GetClient()
	err := cli.ContainerRemove(ctx, container_id, types.ContainerRemoveOptions{})
	return err
}