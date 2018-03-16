package local

import (
	"github.com/docker/docker/api/types"
	"fmt"
)

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