package local

import (
	"golang.org/x/net/context"
)

func KillContainer(container_id string) error {
	ctx := context.Background()
	cli := GetClient()
	// 第三个参数是signal, 是发给container的指令，默认为“KILL"
	err := cli.ContainerKill(ctx, container_id, "KILL")
	return err
}