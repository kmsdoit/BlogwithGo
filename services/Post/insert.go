package services

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmsdoit/blog/models"
)

var post models.Post

func InsertByUserEmail(c *gin.Context) {
	c.Header("Content-type", "application/json")

	body := json.NewDecoder(c.Request.Body).Decode(&post)
	postInsert := dbConn.Debug().Save(&post).Error
	if postInsert == nil && body == nil {
		dbConn.Debug().Save(&post)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   post})
	} else {
		c.JSON(500, map[string]string{
			"message": "fail"})
	}

}
