package services

import (
	"github.com/gin-gonic/gin"
	"github.com/kmsdoit/blog/models"
)

func Register(c *gin.Context) {
	var users = models.User{}
	if err := c.BindJSON(&users); err == nil {
		dbErr := dbConn.Debug().Save(&users).Error
		if dbErr == nil {
			users.Password = GenerateHashPassword(users.Password)
			dbConn.Debug().Save(&users)
			c.JSON(200, gin.H{
				"status": "ok",
				"data":   users,
			})
		} else {
			c.JSON(500, nil)
			return
		}
	}
}
