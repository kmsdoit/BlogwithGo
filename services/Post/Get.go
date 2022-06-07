package services

import (
	"github.com/gin-gonic/gin"
	"github.com/kmsdoit/blog/models"
)

func GetAllPosts(c *gin.Context) {
	var posts = models.GetPosts()

	err := dbConn.Find(&posts).Error
	if err == nil {
		c.JSON(200, gin.H{
			"status": "ok",
			"data":   posts,
		})
	} else {
		c.JSON(500, nil)
		return
	}
}
