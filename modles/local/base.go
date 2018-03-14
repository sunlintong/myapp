package local

import (
	"github.com/docker/docker/client"
	"sync"
)

var (
	globalClient *client.Client
	once sync.Once
)

// 单例模式获取本地docker的Client
func GetClient() *client.Client {
	var err error
	once.Do(func() {
		globalClient, err = client.NewEnvClient()
		if err != nil {
			panic(err)
		}
	})
	return globalClient
}
