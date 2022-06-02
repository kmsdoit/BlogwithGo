package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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
		access_token,err := CreateJWT(user.Email)
		if(err != nil) {
			c.JSON(http.StatusInternalServerError,map[string]string{
				"message" : "server Error Pleaze check your url",
			})
		}else{
			c.SetCookie("access_token",access_token,60*60*24,"/","localhost:8003",false,true)
			c.JSON(http.StatusOK,map[string]string{
					"message" : "토큰 발급 완료",
					"access_token" : access_token,
			})
		}		
	}
}

func VerifyAccessToken(c *gin.Context) {
	ctoken, err := c.Request.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized,gin.H{
			"status" : 401,
			"message" : "Get Cookie failed",
		})
		c.Abort()
		return
	}
	tknstr := ctoken.Value

	if tknstr == ""{
		c.JSON(http.StatusUnauthorized,gin.H{
			"status" : 401,
			"message" : "token is None",
		})

		c.Abort()
		return
	}

	claims := jwt.MapClaims{}

	token,err := jwt.ParseWithClaims(tknstr,&claims,func(token *jwt.Token) (interface{}, error) {
		return []byte(mySignKey),nil
	})

	fmt.Printf("token : %v\n", token)

	if err != nil {
		c.JSON(http.StatusUnauthorized,gin.H{
			"status" : 401,
			"message" : "토큰 인증 실패",
		})
	}else {
		c.JSON(http.StatusOK,gin.H{
			"status" : 200,
			"message" : "토큰 인증 완료",
		})
	}

	return
 }

