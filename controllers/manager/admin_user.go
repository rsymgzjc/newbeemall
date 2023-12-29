package manager

import (
	"errors"
	"newbeemall/controllers"
	"newbeemall/dao/mysql"
	"newbeemall/dao/redis"
	"newbeemall/logic/manager"
	"newbeemall/models"
	"strings"

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

func AdminLoginHandler(c *gin.Context) {
	p := new(models.AdminLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	token, err := manager.AdminLogin(p)
	if err != nil {
		zap.L().Error("logic.AdminLogin with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	err = redis.AddAdminToken(token)
	if err != nil {
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "管理员登录成功")
}

func AdminLogoutHandler(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	splitStr := strings.Split(authHeader, " ")
	Token := splitStr[1]
	if err := redis.DeleteUserToken(Token); err != nil {
		controllers.ResponseError(c, controllers.CodeLogoutFailed)
		return
	}
	controllers.ResponseSuccess(c, "管理员退出成功")
}

func AdminUpdateNameHandler(c *gin.Context) {
	p := new(models.AdminUpdate)
	if err := c.ShouldBindJSON(p); err != nil {

	}
}
