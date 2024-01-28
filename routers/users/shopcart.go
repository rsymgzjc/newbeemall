package users

import (
	"newbeemall/controllers/user"
	"newbeemall/middlewares"

	"github.com/gin-gonic/gin"
)

type MallShopCartRouter struct {
}

func (m *MallUserRouter) InitMallShopCartRouter(r *gin.RouterGroup) {
	r.Use(middlewares.JWTUserAuthMiddleware())
	{
		r.POST("/shop-cart", user.AddShopCartHandler) //添加购物车
		//r.GET("/shop-cart", user.GetShopCartList) //查看购物车清单
	}
}
