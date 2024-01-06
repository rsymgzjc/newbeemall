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
		r.PUT("goods", manager.UpdateGoodsInfoHandler)                   //更新商品信息
		r.GET("goods/:id", manager.GetGoodsInfoByIDHandler)              //根据id获取商品信息
		r.GET("/goods/list", manager.GetGoodsListHandler)                //分页显示商品信息
	}
}
