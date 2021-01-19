package main

import (
	"log"

	"github.com/francobottoni/TwitterApp/database"
	"github.com/francobottoni/TwitterApp/handlers"
)

func main() {
	if database.CheckConection() == false {
		log.Fatal("Missing conection with DB")
		return
	}
	handlers.InitHandler()
}
