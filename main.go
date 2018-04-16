package main

import (
	_ "myapp/routers"

	"github.com/astaxie/beego"
	"myapp/modles/db"
	"time"
	//	"myapp/test"
)


func main() {
	go func() {
		ticker := time.NewTicker(15 * time.Second)
		for range ticker.C {
			db.SyncContainers()
		}
	}()
	beego.Run()
}
