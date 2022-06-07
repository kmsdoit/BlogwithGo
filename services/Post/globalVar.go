package services

import (
	"github.com/jinzhu/gorm"
	"github.com/kmsdoit/blog/models"
)

var dbConn *gorm.DB
var posts = models.GetPosts()
