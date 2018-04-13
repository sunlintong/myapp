package local

import (
	"golang.org/x/net/context"
)

func AddTag(source, target string) error {
	ctx := context.Background()
	cli := GetClient()
	err := cli.ImageTag(ctx, source, target)
	return err
}
