package services

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/gorm"
	"github.com/kmsdoit/blog/models"
	"golang.org/x/crypto/bcrypt"
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

func Register(c *gin.Context) {
	c.Header("Content-type", "application/json")
	var users = models.User{}

	err := dbConn.Debug().Save(&users).Error
	body := json.NewDecoder(c.Request.Body).Decode(&users)
	if err == nil && body == nil{
		users.Password = GenerateHashPassword(users.Password)
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

func Login(c *gin.Context) {
	var user = new(models.User)

	if err := c.Bind(user);  err != nil {
        c.JSON(http.StatusBadRequest, map[string]string{
            "message": "bad request",
        })
    }

	inputpw := user.Password

	result := dbConn.Find(user, "email = ?", user.Email)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, map[string]string{
            "message": "bad request",
        })
	}

	res := ComparePassword(user.Password,inputpw)

	if !res {
		c.JSON(http.StatusUnauthorized,map[string]string{
            "message": "허가하지 않는 계정입니다",
        })
	}else {
		c.JSON(http.StatusOK,map[string]string{
			"messange" : "login success",
		})
		
	}
}

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
	mySignKey := []byte(os.Getenv("SECRET_KEY"))

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