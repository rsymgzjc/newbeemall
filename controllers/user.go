package controllers

import (
	"errors"
	"newbeemall/dao/mysql"
	"newbeemall/logic"
	"newbeemall/models"

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
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//业务逻辑
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp with invalid param", zap.Error(err))
		if errors.Is(err, mysql.UserExist1) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "注册成功")
}

func UserLoginHandler(c *gin.Context) {
	//获取参数和校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//业务处理
	Token, err := logic.Login(p)
	if err != nil {
		if errors.Is(err, mysql.UserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, Token)
}
