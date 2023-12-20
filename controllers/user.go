package controllers

import (
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

		}
	}
}
