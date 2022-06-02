package services

import (
	"github.com/jinzhu/gorm"
	"github.com/kmsdoit/blog/models"
)

func SetDB(db *gorm.DB) {
	dbConn = db
	var post = models.GetPost()
	dbConn.AutoMigrate(&post)
}
