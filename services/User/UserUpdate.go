package services

import "github.com/gin-gonic/gin"

func UserUpdate(c *gin.Context) {

	if err := c.BindJSON(&user); err == nil {
		dbErr := dbConn.Debug().Save(&user).Where(&user.Email).Error
		if dbErr == nil {
			dbConn.Debug().Save(&user).Where(&user.Email)
			c.JSON(200, gin.H{
				"status": "ok",
				"data":   &user})
		} else {
			c.JSON(500, map[string]string{
				"message": "fail"})
		}
	}

}
