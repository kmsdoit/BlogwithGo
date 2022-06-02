package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	"github.com/kmsdoit/blog/models"
	"golang.org/x/crypto/bcrypt"
)

var mySignKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateHashPassword(password string) (string) {
	hash,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err != nil {
		return "" 
	}
	return string(hash)
}

func ComparePassword(hash,password string) (bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false
		}
		return false
	}

	return true
}

func CreateJWT(Email string) (string, error) {
	

	aToken := jwt.New(jwt.SigningMethodHS256)
	claims := aToken.Claims.(jwt.MapClaims)
	claims["Email"] = Email
	claims["exp"] = time.Now().Add(time.Minute * 20).Unix()

	token,err := aToken.SignedString(mySignKey)
	if err != nil {
		return "",err
	}else {
		return token,nil
	}
}


func SetDB(db *gorm.DB) {
	dbConn = db
	var user = models.GetUser()
	dbConn.AutoMigrate(&user)
}