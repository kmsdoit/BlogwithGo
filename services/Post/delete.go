package services

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Delete(c *gin.Context) {
	if err := c.BindJSON(&post); err == nil {
		dbErr := dbConn.Debug().Delete(&post).Error
		if dbErr == nil {
			dbConn.Where("id = ?", &post.ID).Unscoped().Delete(&post)
			log.Print(&post.ID)
			c.JSON(200, gin.H{
				"message": "success",
			})
		} else {
			c.JSON(500, gin.H{
				"message": "fail",
			})
		}
	}
}
