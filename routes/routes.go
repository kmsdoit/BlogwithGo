package routes

import (
	"github.com/gin-gonic/gin"
	post "github.com/kmsdoit/blog/services/Post"
	user "github.com/kmsdoit/blog/services/User"
)

func Router() {
	r := gin.Default()
	r.GET("/users", user.GetAllUsers)
	r.GET("/users/info", user.GetFindByEmail)
	r.POST("/register", user.Register)
	r.POST("/login", user.Login)
	r.GET("/posts", post.GetAllPosts)
	r.GET("/verify", user.VerifyAccessToken)
	r.POST("/post", post.InsertByUserEmail)
	r.Run(":8003")
}