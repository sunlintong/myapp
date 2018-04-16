package main

import (
	_ "myapp/routers"

	"github.com/astaxie/beego"
	"myapp/modles/db"
	//	"myapp/test"
)


func main() {
	go db.SyncContainers()
	beego.Run()
}
