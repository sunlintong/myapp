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

func GetAllContainersGrepIds(ids []string) ([]types.Container, error) {
	c := make([]types.Container, len(ids))
	containers, err := GetAllContainers()
	for i, id := range ids {
		for _, container := range containers {
			if container.ID == id {
				c[i] = container
				break
			}
		}
	}
	return c, err
}

func GetRunningContainersGrepIds(ids []string) ([]types.Container, error) {
	containers, err := GetRunningContainers()
	// 长度为containers与ids的交集，小于等于其中任一个
	// 长度不定，初始为0，逐个append
	c := make([]types.Container, 0, len(ids))
	for _, cc := range containers {
		for _, id := range ids {
			if cc.ID == id {
				c = append(c, cc)
				// 每匹配一个便跳出内循环
				break
			}
		}
	}
	return c, err
}
