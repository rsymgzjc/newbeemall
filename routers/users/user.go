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
	r.Use(middlewares.JWTUserAuthMiddleware())
	{
		r.PUT("/user/info", user.UserUpdateHandler)    //用户信息修改
		r.GET("/user/info", user.UserGetInfoHandler)   //获取用户信息
		r.POST("/user/logout", user.UserLogoutHandler) //用户登出
	}
}
