package manager

import (
	"newbeemall/controllers/manager"
	"newbeemall/middlewares"

	"github.com/gin-gonic/gin"
)

type ManageGoodsInfoRouter struct {
}

func (m *ManageGoodsInfoRouter) InitManageGoodsInfoRouter(r *gin.RouterGroup) {
	r.Use(middlewares.JWTAdminAuthMiddleware())
	{
		r.POST("/goods", manager.CreateGoodsInfoHandler)                 //新建商品信息
		r.PUT("/goods/status/:status", manager.ChangeGoodsStatusHandler) //上下架
	}
}
