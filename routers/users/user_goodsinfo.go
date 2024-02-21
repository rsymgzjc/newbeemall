package users

import (
	"newbeemall/controllers/user"
	"newbeemall/middlewares"

	"github.com/gin-gonic/gin"
)

type MallGoodsInfoIndexRouter struct {
}

func (m *MallGoodsInfoIndexRouter) InitMallGoodsInfoIndexRouter(r *gin.RouterGroup) {
	r.Use(middlewares.JWTUserAuthMiddleware())
	{
		r.GET("/search", user.SearchGoodsHandler) //商品搜索
	}
}
