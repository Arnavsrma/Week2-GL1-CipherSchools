package main

import (
	"log"

	"github.com/Arnavsrma/Week2-GL1-Cipherschool/database"
	"github.com/Arnavsrma/Week2-GL1-Cipherschool/routers"
)

func main() {
	database.Setup()
	engine := router.Router()
	err := engine.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
}
