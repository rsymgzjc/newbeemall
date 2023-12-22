package routers

import (
	"net/http"
	"newbeemall/logger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//商城前端路由
	mallRouter := RouterGroupApp.Mall
	mallGroup := r.Group("api")
	{
		mallRouter.InitMallUserRouter(mallGroup)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
