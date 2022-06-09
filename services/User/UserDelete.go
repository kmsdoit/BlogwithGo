package services

import (
	"github.com/gin-gonic/gin"
	"log"
)

func UserDelete(c *gin.Context) {
	if err := c.BindJSON(&user); err == nil {
		dbErr := dbConn.Debug().Delete(&user).Error
		if dbErr == nil {
			dbConn.Where("email = ?", &user.Email).Unscoped().Delete(&user)
			log.Print(&user.Email)
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
