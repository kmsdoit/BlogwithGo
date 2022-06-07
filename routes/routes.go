package routes

import "github.com/gin-gonic/gin"

var router = gin.Default()

func Router() {

	PostRouter()
	UserRouter()
	err := router.Run(":8003")
	if err != nil {
		return
	}
}
