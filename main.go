package main

import (
	"log"

	"final-task-of-diploma/pkg/db"
	"final-task-of-diploma/pkg/server"
)

func main() {
	if err := db.Init("scheduler.db"); err != nil {
		log.Fatal(err)
	}

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
