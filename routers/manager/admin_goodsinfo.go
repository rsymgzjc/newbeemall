package manager

import (
	"newbeemall/controllers/manager"

	"github.com/gin-gonic/gin"
)

type ManageGoodsInfoRouter struct {
}

func (m *ManageGoodsInfoRouter) InitManageGoodsInfoRouter(r *gin.RouterGroup) {
	{
		r.POST("goods", manager.CreateGoodsInfoHandler) //新建商品信息
	}
}
