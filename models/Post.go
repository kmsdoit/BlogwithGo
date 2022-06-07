package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255)" json:"Title"`
	Content string `gorm:"type:text" json:"Content"`
	Writer  string `gorm:"type:text;not null" json:"Writer"`
}

func GetPost() Post {
	var post Post

	return post
}

func GetPosts() []Post {
	var posts []Post

	return posts
}
