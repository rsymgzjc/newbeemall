package manager

import (
	"github.com/gin-gonic/gin"
	"newbeemall/controllers/manager"
)

type ManageAdminUserRouter struct {
}

func (m *ManageAdminUserRouter) InitManageAdminUserRouter(r *gin.RouterGroup) {
	r.POST("/admin/login", manager.AdminLoginHandler)
}
