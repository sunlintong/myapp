package local

import (
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)


// 根据container_id启动一个容器
// 并不判断容器状态
// 对应命令docker start <container_id>
func StartContainer(container_id string) error {
	ctx := context.Background()
	cli := GetClient()
	err := cli.ContainerStart(ctx, container_id, types.ContainerStartOptions{})
	return err
}
