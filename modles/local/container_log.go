package local

import (
	"golang.org/x/net/context"
	"github.com/docker/docker/api/types"
	"io"
)

func GetContainerLog(container_id string) (io.ReadCloser, error) {
	ctx := context.Background()
	cli := GetClient()
	// 设置为返回stdout的输出流
	options := types.ContainerLogsOptions{}
	options.ShowStdout = true
	options.ShowStderr = true
//	options.Details = true
//	options.Follow = true
	options.Timestamps = true
	out, err := cli.ContainerLogs(ctx, container_id, options)
	return out, err
}