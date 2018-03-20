package local

import (
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"io"
)


// 拉镜像到本地
// 注意返回了一个io.ReadCloser接口，其中有pull image时的输出流
// 对应命令docker pull <image_name>
func PullImage(imageName string, ims ...string ) (io.ReadCloser, error) {
	cli := GetClient()
	out, err := cli.ImagePull(context.Background(), imageName, types.ImagePullOptions{})
//	defer out.Close()
	return out, err
}
