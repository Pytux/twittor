package main

import (
	"log"

	"github.com/pytux/twittor/db"
	"github.com/pytux/twittor/handlers"
)

func main() {

	if !db.CheckConnection() {
		log.Fatal("Error, cant connect to DB.")
		return
	}

	handlers.Handlers()

}
