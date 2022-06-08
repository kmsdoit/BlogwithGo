package routes

import (
	post "github.com/kmsdoit/blog/services/Post"
)

func PostRouter() {
	postApi := router.Group("/api/post")
	{
		postApi.GET("/getAllPost", post.GetAllPosts)
		postApi.POST("/post", post.InsertByUserEmail)
		postApi.PATCH("/postUpdate", post.Update)
		postApi.DELETE("/postDelete", post.Delete)
	}
}
