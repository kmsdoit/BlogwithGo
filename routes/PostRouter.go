package routes

import (
	post "github.com/kmsdoit/blog/services/Post"
)

func PostRouter() {
	postApi := router.Group("/api/post")
	{
		postApi.GET("/getAllPost", post.GetAllPosts)
		postApi.POST("/post", post.InsertByUserEmail)
		postApi.POST("/postUpdate", post.Update)
		postApi.POST("/postDelete", post.Delete)
	}
}
