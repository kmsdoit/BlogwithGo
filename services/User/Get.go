package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmsdoit/blog/models"
)



func GetAllUsers(c *gin.Context) {	
	var users = models.GetUsers()

	err := dbConn.Find(&users).Error
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   users,
		})
	}else {
		c.JSON(http.StatusNoContent, nil)
		return
	}
}

func GetFindByEmail(c *gin.Context) {
	var email =  c.Request.URL.Query().Get("email") 
	var user = models.User{}

	err := dbConn.Where("email =?",email).Find(&user).Error

	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":  user,
		})
	}else {
		c.JSON(http.StatusInternalServerError,nil)
		return
	}
}