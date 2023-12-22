package user

import (
	"errors"
	"newbeemall/controllers"
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
			controllers.ResponseError(c, controllers.CodeInvalidParam)
			return
		}
		controllers.ResponseErrorWithMsg(c, controllers.CodeInvalidParam, controllers.RemoveTopStruct(errs.Translate(controllers.Trans)))
		return
	}
	//业务逻辑
	if err := logic.SignUp(p); err != nil {
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
	Token, err := logic.Login(p)
	if err != nil {
		if errors.Is(err, mysql.UserNotExist) {
			controllers.ResponseError(c, controllers.CodeUserNotExist)
			return
		}
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, Token)
}
