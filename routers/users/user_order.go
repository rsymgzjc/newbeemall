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
	}
}
