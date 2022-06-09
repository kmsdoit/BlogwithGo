package routes

import (
	user "github.com/kmsdoit/blog/services/User"
)

func UserRouter() {
	userApi := router.Group("/api/user")
	{
		userApi.GET("/", user.GetAllUsers)
		userApi.GET("/info", user.GetFindByEmail)
		userApi.POST("/register", user.Register)
		userApi.POST("/login", user.Login)
		userApi.GET("/verify", user.VerifyAccessToken)
		userApi.PATCH("/userUpdate", user.UserUpdate)
		userApi.DELETE("/userDelete", user.UserDelete)
	}

}
