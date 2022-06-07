package services

import (
	"github.com/gin-gonic/gin"
)

func GetAllPosts(c *gin.Context) {
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
