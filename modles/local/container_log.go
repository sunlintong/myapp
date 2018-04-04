package local

import (
	"golang.org/x/net/context"
	"github.com/docker/docker/api/types"
	"io"
)

func GetContainerLog(container_id string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
	ctx := context.Background()
	cli := GetClient()
	// 设置为返回stdout的输出流
	out, err := cli.ContainerLogs(ctx, container_id, options)
	return out, err
}