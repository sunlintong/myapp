package local

import (
//	"github.com/docker/docker/api/types/container"
	"golang.org/x/net/context"
)


// 根据container_id停止一个容器
// 对应命令docker stop
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

