package users

import (
	"newbeemall/controllers/user"
	"newbeemall/middlewares"

	"github.com/gin-gonic/gin"
)

type MallUserAddressRouter struct {
}

func (m *MallUserRouter) InitMallUserAddressRouter(r *gin.RouterGroup) {
	r.Use(middlewares.JWTUserAuthMiddleware())
	{
		r.GET("/address", user.GetAddressListHandler)           //用户地址
		r.POST("/address", user.AddressAddHandler)              //添加地址
		r.PUT("/address", user.UpdateAddressHandler)            //修改用户地址
		r.GET("/address/:addressid", user.GetAddrDetailHandler) //获取地址详情
		r.GET("/address/default", user.GetDefaultAddrHandler)   //获取默认地址
		r.DELETE("/address/:addressid", user.DelAddrHandler)    //删除地址
	}
}
