package manager

import (
	"newbeemall/controllers/manager"
	"newbeemall/middlewares"

	"github.com/gin-gonic/gin"
)

type ManageIndexConfigRouter struct {
}

func (m *ManageIndexConfigRouter) InitManageIndexConfigRouter(r *gin.RouterGroup) {
	r.Use(middlewares.JWTAdminAuthMiddleware())
	{
		r.POST("indexconfig", manager.CreateIndexConfigHandler) //新建首页配置项
	}
}
