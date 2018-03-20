package local

import (
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)


// 获取在运行的容器列表
// 对应命令docker ps
func GetRunningContainers() ([]types.Container, error) {
	cli := GetClient()
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	return containers, err
}

// 获取所有容器，包括未运行的
func GetAllContainers() ([]types.Container, error) {
	cli := GetClient()	
	options := types.ContainerListOptions{}
	options.All = true
	containers, err := cli.ContainerList(context.Background(), options)
	return containers, err
}