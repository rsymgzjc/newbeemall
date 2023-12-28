package manager

import (
	"errors"
	"newbeemall/controllers"
	"newbeemall/dao/mysql"
	"newbeemall/logic/manager"
	"newbeemall/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AdminSignupHandler(c *gin.Context) {
	p := new(models.AdminSignup)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	if err := manager.AdminSignup(p); err != nil {
		zap.L().Error("logic.SignUp with invalid param", zap.Error(err))
		if errors.Is(err, mysql.AdminExist1) {
			controllers.ResponseError(c, controllers.CodeAdminExist)
			return
		}
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "注册成功")
}
