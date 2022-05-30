package main

import (
	"log"

	"github.com/kmsdoit/blog/routes"
	"github.com/kmsdoit/blog/services"
	"github.com/kmsdoit/blog/utility"
)

func main() {
	var db = utility.GetConnection()
	services.SetDB(db)
	log.Println("Listening on Port 8000")
	routes.UserRouter()
}
