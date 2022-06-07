package services

import (
	"github.com/gin-gonic/gin"
	"github.com/kmsdoit/blog/models"
)

var post models.Post

func InsertByUserEmail(c *gin.Context) {
	if err := c.BindJSON(&post); err == nil {
		postInsert := dbConn.Debug().Save(&post).Error
		if postInsert == nil {
			dbConn.Debug().Save(&post)
			c.JSON(200, gin.H{
				"status": "ok",
				"data":   post})
		} else {
			c.JSON(500, map[string]string{
				"message": "fail"})
		}
	}
}
