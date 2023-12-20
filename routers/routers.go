package routers

import (
	"net/http"
	"newbeemall/controllers"
	"newbeemall/logger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//商城注册
	r.POST("/user/signup", controllers.UserSignUpHandler)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
