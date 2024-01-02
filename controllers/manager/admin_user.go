package manager

import (
	"errors"
	"newbeemall/controllers"
	"newbeemall/dao/mysql"
	"newbeemall/dao/redis"
	"newbeemall/logic/manager"
	"newbeemall/models"
	"path/filepath"
	"strconv"
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
	controllers.ResponseSuccess(c, token)
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
		zap.L().Error("Update with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	adminid, err := controllers.GetCurrentUser(c)
	if err != nil {
		controllers.ResponseError(c, controllers.CodeNeedLogin)
		return
	}
	if err = manager.AdminUpdateName(p, adminid); err != nil {
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "更新名字成功")
}

func AdminUpdatePasswprd(c *gin.Context) {
	p := new(models.AdminUpdate)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Update with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	adminid, err := controllers.GetCurrentUser(c)
	if err != nil {
		controllers.ResponseError(c, controllers.CodeNeedLogin)
		return
	}
	if err = manager.AdminUpdatePassword(p, adminid); err != nil {
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "更新密码成功")
}

func GetUsersListHandler(c *gin.Context) {
	page, size := controllers.GetPageInfo(c)
	data, err := manager.GetUsersList(page, size)
	if err != nil {
		zap.L().Error("logic.GetUserList with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, data)
}

func LockUserHandler(c *gin.Context) {
	locks := c.Param("lockstatus")
	lockstatus, err := strconv.ParseInt(locks, 10, 64)
	if err != nil {
		return
	}
	p := new(models.UserIds)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Lock with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	if err := manager.LockUsers(p, lockstatus); err != nil {
		zap.L().Error("manager.LockUsers failed", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "操作成功")
}

func GetAdminDetailHandler(c *gin.Context) {
	adminid, err := controllers.GetCurrentUser(c)
	if err != nil {
		controllers.ResponseError(c, controllers.CodeNeedLogin)
		return
	}
	data, err := manager.GetAdminDetail(adminid)
	if err != nil {
		zap.L().Error("manager.GetAdminDetail failed", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, data)
}

func UploadFileHandler(c *gin.Context) {
	//获取文件
	file, err := c.FormFile("upload")
	if err != nil {
		zap.L().Error("File with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	dst := filepath.Join("A:", file.Filename)
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		zap.L().Error("File with invalid param", zap.Error(err))
		return
	}
	controllers.ResponseSuccess(c, dst)
}
