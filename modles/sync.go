package modles

import (
	"myapp/modles/db"
	"log"
)

func StartSync() {
	log.Println("startsync0")
	go db.SyncContainers()
	go db.SyncImages()
	log.Println("startsync1")
}