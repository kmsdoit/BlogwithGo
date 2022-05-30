package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kmsdoit/blog/services"
)

func UserRouter() {
	r := gin.Default()
	r.GET("/hello", services.GetAllUsers)
	r.GET("/users/info", services.GetFindByEmail)
	r.POST("/users", services.PostInserData)
	r.Run()
}