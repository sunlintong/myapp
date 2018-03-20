package local

import (
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

// 获取镜像列表
// 对应命令docker images
func GetImages() ([]types.ImageSummary, error) {
	cli := GetClient()
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	return images, err
}
