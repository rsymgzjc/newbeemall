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
		r.POST("indexconfig", manager.CreateIndexConfigHandler)   //新建首页配置项
		r.DELETE("indexconfig", manager.DeleteIndexConfigHandler) //删除首页配置项
		r.PUT("indexconfig", manager.UpdateIndexConfigHandler)    //更新首页配置项
		r.GET("indexconfig/:id", manager.GetIndexConfigByID)      //根据id获取首页配置项
		r.GET("indexconfig", manager.GetIndexConfigList)          //获取首页配置项列表
	}
}
