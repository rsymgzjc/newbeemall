package manager

import (
	"newbeemall/controllers/manager"
	"newbeemall/middlewares"

	"github.com/gin-gonic/gin"
)

type ManageAdminUserRouter struct {
}

func (m *ManageAdminUserRouter) InitManageAdminUserRouter(r *gin.RouterGroup) {

	r.POST("/admin/login", manager.AdminLoginHandler) //管理员登陆
	r.Use(middlewares.JWTAdminAuthMiddleware())
	{
		r.POST("/admin/signup", manager.AdminSignupHandler)  //管理员注册
		r.POST("/admin/Logout", manager.AdminLogoutHandler)  //管理员登出
		r.PUT("/admin/name", manager.AdminUpdateNameHandler) //修改名字
	}
}
