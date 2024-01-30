package users

import (
	"github.com/gin-gonic/gin"
	"newbeemall/middlewares"
)

type MallOrderRouter struct {
}

func (m *MallOrderRouter) InitMallOrderRouter(r *gin.RouterGroup) {
	r.Use(middlewares.JWTUserAuthMiddleware())
	{

	}
}
