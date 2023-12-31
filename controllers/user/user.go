package user

import (
	"errors"
	"newbeemall/controllers"
	"newbeemall/dao/mysql"
	"newbeemall/dao/redis"
	"newbeemall/logic/user"
	"newbeemall/models"
	"strings"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func UserSignUpHandler(c *gin.Context) {
	//获取参数和校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			controllers.ResponseError(c, controllers.CodeInvalidParam)
			return
		}
		controllers.ResponseErrorWithMsg(c, controllers.CodeInvalidParam, controllers.RemoveTopStruct(errs.Translate(controllers.Trans)))
		return
	}
	//业务逻辑
	if err := user.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp with invalid param", zap.Error(err))
		if errors.Is(err, mysql.UserExist1) {
			controllers.ResponseError(c, controllers.CodeUserExist)
			return
		}
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "注册成功")
}

func UserLoginHandler(c *gin.Context) {
	//获取参数和校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	//业务处理
	Token, err := user.Login(p)
	if err != nil {
		if errors.Is(err, mysql.UserNotExist) {
			controllers.ResponseError(c, controllers.CodeUserNotExist)
			return
		}
		controllers.ResponseError(c, controllers.CodeUserLocked)
		return
	}
	//添加Token到redis
	if err := redis.AddUserToken(Token); err != nil {
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, Token)
}

func UserUpdateHandler(c *gin.Context) {
	//获取参数与校验
	p := new(models.ParamUpdate)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Update with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	//获取用户ID
	userid, err := controllers.GetCurrentUser(c)
	if err != nil {
		zap.L().Error("获取用户ID失败", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeNeedLogin)
		return
	}
	//业务处理
	if err := user.Update(p, userid); err != nil {
		zap.L().Error("logic.Update with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "更新成功")
}

func UserGetInfoHandler(c *gin.Context) {
	//获取用户ID
	userid, err := controllers.GetCurrentUser(c)
	if err != nil {
		zap.L().Error("获取用户ID失败", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeNeedLogin)
		return
	}
	//业务处理
	data, err := user.GetInfo(userid)
	if err != nil {
		zap.L().Error("logic.GetInfo failed", zap.Error(err))
		return
	}
	controllers.ResponseSuccess(c, data)
}

func UserLogoutHandler(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	splitStr := strings.Split(authHeader, " ")
	Token := splitStr[1]
	if err := redis.DeleteUserToken(Token); err != nil {
		controllers.ResponseError(c, controllers.CodeLogoutFailed)
		return
	}
	controllers.ResponseSuccess(c, "退出成功")
}
