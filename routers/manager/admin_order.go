package manager

import (
	"newbeemall/controllers/manager"
	"newbeemall/middlewares"

	"github.com/gin-gonic/gin"
)

type ManageOrderRouter struct {
}

func (m *ManageOrderRouter) InitManageOrderRouter(r *gin.RouterGroup) {
	r.Use(middlewares.JWTAdminAuthMiddleware())
	{
		r.PUT("orders/checkdone", manager.CheckDoneOrderHandler) //发货
	}
}
