package services

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/kmsdoit/blog/models"
)

var dbConn *gorm.DB

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

func PostInserData(c *gin.Context) {
	c.Header("Content-type", "application/json")
	var users = models.User{}

	err := dbConn.Debug().Save(&users).Error
	body := json.NewDecoder(c.Request.Body).Decode(&users)

	if err == nil && body == nil{
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

func SetDB(db *gorm.DB) {
	dbConn = db
	var user = models.GetUser()
	dbConn.AutoMigrate(&user)
}