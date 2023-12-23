package users

import (
	"newbeemall/controllers/user"
	"newbeemall/middlewares"

	"github.com/gin-gonic/gin"
)

type MallUserRouter struct {
}

func (m *MallUserRouter) InitMallUserRouter(r *gin.RouterGroup) {
	{
		r.POST("/user/signup", user.UserSignUpHandler) //用户注册
		r.POST("/user/login", user.UserLoginHandler)   //用户登录
	}
	r.Use(middlewares.JWTAuthMiddleware())
	{
		r.PUT("/user/update", user.UserUpdateHandler) //用户信息修改
	}
}
