package users

import (
	"newbeemall/controllers/user"
	"newbeemall/middlewares"

	"github.com/gin-gonic/gin"
)

type MallGoodsCategoryIndexRouter struct {
}

func (m *MallGoodsInfoIndexRouter) InitMallGoodsCategoryIndexRouter(r *gin.RouterGroup) {
	r.Use(middlewares.JWTUserAuthMiddleware())
	{
		r.GET("categories", user.GetGoodsCategoryHandler) //获取分类数据
	}
}
