package services

import (
	"github.com/gin-gonic/gin"
	"github.com/kmsdoit/blog/models"
)

func GetAllUsers(c *gin.Context) {
	var users = models.GetUsers()

	err := dbConn.Find(&users).Error
	if err == nil {
		c.JSON(200, gin.H{
			"status": "ok",
			"data":   users,
		})
	} else {
		c.JSON(400, nil)
		return
	}
}

func GetFindByEmail(c *gin.Context) {
	var email = c.Request.URL.Query().Get("email")
	var user = models.User{}

	err := dbConn.Where("email =?", email).Find(&user).Error

	if err == nil {
		c.JSON(200, gin.H{
			"status": "ok",
			"data":   user,
		})
	} else {
		c.JSON(500, nil)
		return
	}
}
