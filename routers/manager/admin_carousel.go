package manager

import (
	"newbeemall/controllers/manager"
	"newbeemall/middlewares"

	"github.com/gin-gonic/gin"
)

type ManageCarouselRouter struct {
}

func (m *ManageCarouselRouter) InitManageCarouselRouter(r *gin.RouterGroup) {
	r.Use(middlewares.JWTAdminAuthMiddleware())
	{
		r.POST("carousels", manager.CreateCarouselHandler) //新建轮播图
	}
}
