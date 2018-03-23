package local

import (
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

func RemoveImage(image_id string) ([]types.ImageDeleteResponseItem, error){
	cli := GetClient()
	ctx := context.Background()
	items, err := cli.ImageRemove(ctx, image_id, types.ImageRemoveOptions{})
	return items, err
}