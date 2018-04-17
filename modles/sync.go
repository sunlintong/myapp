package modles

import (
	"myapp/modles/db"
)

func StartSync() {
	go db.SyncContainers()
	go db.SyncImages()
}