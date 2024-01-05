package manager

import (
	"newbeemall/controllers/manager"
	"newbeemall/middlewares"

	"github.com/gin-gonic/gin"
)

type ManageGoodsCategoryRouter struct {
}

func (m *ManageGoodsCategoryRouter) InitManageGoodsCategoryRouter(r *gin.RouterGroup) {
	r.Use(middlewares.JWTAdminAuthMiddleware())
	{
		r.POST("/categories", manager.CreateCategoryHandler) //创建商品类别
		r.PUT("/categories", manager.UpdateCategoryHandler)  //更新商品类别
		r.GET("/categories", manager.GetCategoryListHandler) //获取分类信息
		r.GET("/categories/:id", manager.GetCategoryHandler) //根据id获取分类
		r.DELETE("/categories", manager.DelCategoryHandler)  //删除指定的分类

	}
}
