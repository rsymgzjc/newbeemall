package manager

import (
	"newbeemall/controllers/manager"

	"github.com/gin-gonic/gin"
)

type ManageAdminUserRouter struct {
}

func (m *ManageAdminUserRouter) InitManageAdminUserRouter(r *gin.RouterGroup) {
	{
		r.POST("/admin/signup", manager.AdminSignupHandler) //管理员注册
		//r.POST("/admin/login", manager.AdminLoginHandler)   //管理员登陆
	}

}
