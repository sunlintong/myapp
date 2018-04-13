package local

import (
	"github.com/docker/docker/api/types/container"
	"golang.org/x/net/context"
)

// 通过配置结构体config创建一个容器
func CreateContainer(image_name string, cmd []string) (container.ContainerCreateCreatedBody, error) {
	ctx := context.Background()
	cli := GetClient()
	config := new(container.Config)
	config.Image = image_name
	config.Cmd = cmd
	resp, err := cli.ContainerCreate(ctx, config, nil, nil, "")
	return resp, err
}
