package main

import (
	"log"
	"time"

	"github.com/yunuscvlk/shellix-task/internal/db"
)

func main() {
	d := db.Initialize()

	log.Println("Worker: started!")
	
	for {
		d.DeleteOneDayOld()

		log.Println("Data Control: Sentences older than 1 day has been deleted.")

		time.Sleep(6 * time.Hour)
	}
}