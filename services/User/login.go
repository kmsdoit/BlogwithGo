package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmsdoit/blog/models"
)

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

