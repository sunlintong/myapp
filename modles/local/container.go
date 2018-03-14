package local

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"golang.org/x/net/context"
	
)


// 获取在运行的容器列表
// 对应命令docker ps
func GetRunningContainers() ([]types.Container, error) {
	cli := GetClient()
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	return containers, err
}

// 通过配置结构体config创建一个容器
func CreateContainer(config *container.Config)(container.ContainerCreateCreatedBody, error) {
	cli := GetClient()
	resp, err := cli.ContainerCreate(context.Background(), config, nil, nil, "")
	return resp, err
}


// 根据container_id启动一个容器
// 并不判断容器状态
// 对应命令docker start <container_id>
func StartContainer(container_id string) error {
	ctx := context.Background()
	cli := GetClient()
	err := cli.ContainerStart(ctx, container_id, types.ContainerStartOptions{})
	return err
}


// 根据container_id停止一个容器
// 对应命令docker dtop
func StopContainer(container_id string) error {
	cli := GetClient()
	err := cli.ContainerStop(context.Background(), container_id, nil)
	return err
}


// 停止所有运行中的容器
func StopAllContainers() error{
	containers, err := GetRunningContainers()
	for _, container := range containers {
		err1 := StopContainer(container.ID)
		if err1 != nil {
			return err1
		} 
	}
	return err
}