package main

import (
	_ "myapp/routers"

	"github.com/astaxie/beego"
	//	"myapp/test"
)

func main() {
	//	test.TestGetImages()
	//	test.TestGetRunningContainers()
	//	test.TestPullImage()
	//	test.TestStartContainer()
	//	test.TestStopContainer()
	//	test.TestStopAllContainers()
	
	//	StaticDir[] = "assets"
	beego.Run()
}
