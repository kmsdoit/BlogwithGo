package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmsdoit/blog/models"
)

func GetAllPosts(c *gin.Context) {
	var posts = models.GetPosts()

	err := dbConn.Find(&posts).Error
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   posts,
		})
	} else {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
}
