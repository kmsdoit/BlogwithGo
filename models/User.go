package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(255)" json:"Email"`
	Password string `gorm:"type:text" json:"password"`
	Name     string `gorm:"type:text" json:"name"`
}

func GetUser() User {
	var user User

	return user
}

func GetUsers() []User {
	var users []User

	return users
}
