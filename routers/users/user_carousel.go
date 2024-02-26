package users

import (
	"newbeemall/controllers/user"
	"newbeemall/middlewares"

	"github.com/gin-gonic/gin"
)

type MallCarouselIndexRouter struct {
}

func (m *MallCarouselIndexRouter) InitMallCarouselIndexRouter(r *gin.RouterGroup) {
	r.Use(middlewares.JWTUserAuthMiddleware())
	{
		r.GET("index-infos", user.GetIndexInfoHandler) //获取首页数据
	}
}
