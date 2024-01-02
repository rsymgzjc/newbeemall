package manager

import (
	"github.com/gin-gonic/gin"
	"newbeemall/controllers/manager"
	"newbeemall/middlewares"
)

type ManageGoodsCategoryRouter struct {
}

func (m *ManageGoodsCategoryRouter) InitManageGoodsCategoryRouter(r *gin.RouterGroup) {
	r.Use(middlewares.JWTAdminAuthMiddleware())
	{
		r.POST("/categories", manager.CreateCategoryHandler) //创建商品类别
		r.PUT("/categories", manager.UpdateCategoryHandler)  //更新商品类别
	}
}
