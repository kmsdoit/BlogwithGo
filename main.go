package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kmsdoit/blog/routes"
	post "github.com/kmsdoit/blog/services/Post"
	user "github.com/kmsdoit/blog/services/User"
	"github.com/kmsdoit/blog/utility"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var db = utility.GetConnection()
	post.SetDB(db)
	user.SetDB(db)
	log.Println("Listening on Port 8000")
	routes.Router()
}
