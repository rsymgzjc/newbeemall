package users

import (
	"newbeemall/controllers/user"
	"newbeemall/middlewares"

	"github.com/gin-gonic/gin"
)

type MallUserAddressRouter struct {
}

func (m *MallUserRouter) InitMallUserAddressRouter(r *gin.RouterGroup) {
	r.Use(middlewares.JWTAuthMiddleware())
	{
		r.POST("/address", user.AddressAddHandler) //添加地址
	}
}
