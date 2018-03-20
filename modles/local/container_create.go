package local

import (
	"github.com/docker/docker/api/types/container"
	"golang.org/x/net/context"
)


// 通过配置结构体config创建一个容器
func CreateContainer(config *container.Config)(container.ContainerCreateCreatedBody, error) {
	cli := GetClient()
	resp, err := cli.ContainerCreate(context.Background(), config, nil, nil, "")
	return resp, err
}