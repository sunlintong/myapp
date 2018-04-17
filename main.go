package main

import (
	_ "myapp/routers"

	"github.com/astaxie/beego"
	"myapp/modles"
)


func main() {
	modles.StartSync()
	beego.Run()
}
