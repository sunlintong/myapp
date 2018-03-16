package local

import (
	"github.com/docker/docker/api/types"
	"fmt"
)

func GetContainerPort(c types.Container) string {
	var pstr string
	for _, port := range c.Ports {
		if &port == nil {
			return "未添加"
		}
		pstr += port.IP + ":" +
				fmt.Sprint(port.PrivatePort) + "->" +
				fmt.Sprint(port.PublicPort) + "/" +
				port.Type + 
				"	"
	}
	return pstr
}