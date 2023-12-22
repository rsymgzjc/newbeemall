package routers

import (
	"net/http"
	"newbeemall/controllers/user"
	"newbeemall/logger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	{
		//商城注册
		r.POST("/user/signup", user.UserSignUpHandler)
		//商城登录
		r.POST("/user/login", user.UserLoginHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
