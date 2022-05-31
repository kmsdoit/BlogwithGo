package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kmsdoit/blog/routes"
	"github.com/kmsdoit/blog/services"
	"github.com/kmsdoit/blog/utility"
)

func main() {
	err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	var db = utility.GetConnection()
	services.SetDB(db)
	log.Println("Listening on Port 8000")
	routes.UserRouter()
}
