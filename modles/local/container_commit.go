package local

import (
	"golang.org/x/net/context"
	"github.com/docker/docker/api/types"
)

func CommitContainer(container_id string, image_name string) (string, error) {
	cli := GetClient()
	ctx := context.Background()
	options := types.ContainerCommitOptions{}
	options.Reference = image_name
	resp, err := cli.ContainerCommit(ctx, container_id, options)
	return resp.ID, err
}