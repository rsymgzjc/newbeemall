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
		r.PUT("orders/checkout", manager.CheckOutOrderHandler)   //出库
		r.PUT("orders/close", manager.CloseOrderHandler)         //商家关闭订单
		r.GET("orders/:orderId", manager.GetOrderHandler)        //根据ID获取Order
		r.GET("orders", manager.GetOrderListHandler)             //获取订单列表
	}
}
