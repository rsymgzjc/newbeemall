package users

import (
	"github.com/gin-gonic/gin"
	"newbeemall/middlewares"
)

type MallShopCartRouter struct {
}

func (m *MallUserRouter) InitMallShopCartRouter(r *gin.RouterGroup) {
	r.Use(middlewares.JWTUserAuthMiddleware())
	{

	}
}
