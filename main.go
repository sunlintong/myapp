package main

import (
	"github.com/astaxie/beego"
	_ "myapp/routers"
	//	"myapp/test"
)

func main() {
	//	test.TestGetImages()
	//	test.TestGetRunningContainers()
	//	test.TestPullImage()
	//	test.TestStartContainer()
	//	test.TestStopContainer()
	//	test.TestStopAllContainers()
	beego.Run()
}
