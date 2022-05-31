package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kmsdoit/blog/services"
)

func UserRouter() {
	r := gin.Default()
	r.GET("/users", services.GetAllUsers)
	r.GET("/users/info", services.GetFindByEmail)
	r.POST("/register", services.Register)
	r.POST("/login", services.Login)
	r.Run()
}