package routes

import (
	"github.com/gin-gonic/gin"
	user "github.com/kmsdoit/blog/services/User"
)

func UserRouter() {
	r := gin.Default()
	r.GET("/users", user.GetAllUsers)
	r.GET("/users/info", user.GetFindByEmail)
	r.POST("/register", user.Register)
	r.POST("/login", user.Login)
	r.Run()
}