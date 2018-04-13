package local

import (
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

// 默认的删除方式，不强制删除，不删除子镜像层
func RemoveImage(image_id string) ([]types.ImageDeleteResponseItem, error){
	cli := GetClient()
	ctx := context.Background()
	// 默认非强制
	items, err := cli.ImageRemove(ctx, image_id, types.ImageRemoveOptions{})
	return items, err
}

// 强制删除，不删除子镜像层
func ForceRemoveImage(image_id string) ([]types.ImageDeleteResponseItem, error) {
	cli := GetClient()
	ctx := context.Background()
	options := types.ImageRemoveOptions{}
	options.Force = true
	options.PruneChildren = true
	items, err := cli.ImageRemove(ctx, image_id, options)
	return items, err
}

