package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Update(c *gin.Context) {
	c.Header("Content-type", "application/json")

	if err := c.BindJSON(&post); err == nil {
		dbErr := dbConn.Debug().Save(&post).Where(&post.ID).Error
		if dbErr == nil {
			dbConn.Debug().Save(&post).Where(&post.ID)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"data":   &post})
		} else {
			c.JSON(500, map[string]string{
				"message": "fail"})
		}
	}

}
