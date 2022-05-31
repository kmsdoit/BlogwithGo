package services

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmsdoit/blog/models"
)

func Register(c *gin.Context) {
	c.Header("Content-type", "application/json")
	var users = models.User{}

	err := dbConn.Debug().Save(&users).Error
	body := json.NewDecoder(c.Request.Body).Decode(&users)
	if err == nil && body == nil{
		users.Password = GenerateHashPassword(users.Password)
		dbConn.Debug().Save(&users)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   users,
		})
	}else {
		c.JSON(http.StatusInternalServerError,nil)
		return
	}
}