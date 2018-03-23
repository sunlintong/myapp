package local

import (
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

func RemoveImage(image_id string) ([]types.ImageDeleteResponseItem, error){
	cli := GetClient()
	ctx := context.Background()
	// 默认非强制
	items, err := cli.ImageRemove(ctx, image_id, types.ImageRemoveOptions{})
	return items, err
}

func ForceRemoveImage(image_id string) ([]types.ImageDeleteResponseItem, error) {
	cli := GetClient()
	ctx := context.Background()
	options := types.ImageRemoveOptions{}
	options.Force = true
	items, err := cli.ImageRemove(ctx, image_id, types.ImageRemoveOptions{})
	return items, err
}