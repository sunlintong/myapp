package local

import (
	"golang.org/x/net/context"
	"github.com/docker/docker/api/types"
	"io"
)

func GetContainerLog(container_id string) (io.ReadCloser, error) {
	ctx := context.Background()
	cli := GetClient()
	out, err := cli.ContainerLogs(ctx, container_id, types.ContainerLogsOptions{})
	return out, err
}