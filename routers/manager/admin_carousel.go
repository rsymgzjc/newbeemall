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
		r.POST("carousels", manager.CreateCarouselHandler)     //新建轮播图
		r.DELETE("carousels", manager.DeleteCarouselHandler)   //删除轮播图
		r.PUT("carousels", manager.UpdateCarouselHandler)      //更新轮播图
		r.GET("carousels/:id", manager.GetCarouselByIDHandler) //根据轮播图id获取
		r.GET("carousels", manager.GetCarouselsListHandler)    //获取轮播图列表
	}
}
