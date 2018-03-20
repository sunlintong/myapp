package local

import (
	"github.com/docker/docker/api/types"
//	"github.com/docker/docker/api/types/container"
	"golang.org/x/net/context"
	"fmt"
)

// 格式化获取容器端口映射字符串
func GetContainerPort(c types.Container) string {
	var pstr string
	for _, port := range c.Ports {
		pstr += port.IP + ":" +
				fmt.Sprint(port.PrivatePort) + "->" +
				fmt.Sprint(port.PublicPort) + "/" +
				port.Type + 
				"	"
	}
	if pstr == "" {
		pstr = "未添加"
	}
	return pstr
}

// 根据container_id停止一个容器
// 对应命令docker stop
func StopContainer(container_id string) error {
	cli := GetClient()
	err := cli.ContainerStop(context.Background(), container_id, nil)
	return err
}
