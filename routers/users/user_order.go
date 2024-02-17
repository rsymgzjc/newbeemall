package users

import (
	"newbeemall/controllers/user"
	"newbeemall/middlewares"

	"github.com/gin-gonic/gin"
)

type MallOrderRouter struct {
}

func (m *MallOrderRouter) InitMallOrderRouter(r *gin.RouterGroup) {
	r.Use(middlewares.JWTUserAuthMiddleware())
	{
		r.POST("/saveorder", user.SaveOrderHandler)               //生成订单
		r.GET("/paysuccess", user.PaySuccessHandler)              //支付成功
		r.PUT("/order/:ordernum/finish", user.FinishOrderHandler) //确认收货接口
		r.PUT("/order/:ordernum/cancel", user.CancelOrderHandler) //取消订单接口
		r.GET("/order/:ordernum", user.OrderDetailHandler)        //订单详情接口
		r.GET("/order", user.GetOrderListHandler)                 //订单列表接口
	}
}
