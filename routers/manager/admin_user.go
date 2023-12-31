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
		r.POST("/admin/signup", manager.AdminSignupHandler)        //管理员注册
		r.POST("/admin/Logout", manager.AdminLogoutHandler)        //管理员登出
		r.PUT("/admin/name", manager.AdminUpdateNameHandler)       //修改名字
		r.PUT("/admin/password", manager.AdminUpdatePasswprd)      //修改密码
		r.GET("/admin/users", manager.GetUsersListHandler)         //分页展示用户信息
		r.PUT("/admin/users/:lockstatus", manager.LockUserHandler) //封锁用户
		r.GET("/admin/profile", manager.GetAdminDetailHandler)     //获取管理员信息
		r.POST("/upload/file", manager.UploadFileHandler)          //上传文件
	}
}
