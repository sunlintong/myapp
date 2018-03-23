package local

import (
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

// 获取镜像列表, 不获取镜像相关操作
// 对应命令docker images
func GetImages() ([]types.ImageSummary, error) {
	cli := GetClient()
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	return images, err
}

func GetImage(image_id string) (types.ImageSummary, error) {
	images, err := GetImages()
	for _, image := range images {
		if image.ID == image_id {
		return image, err
		}
	}
	return types.ImageSummary{}, err
}
