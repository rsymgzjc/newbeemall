package user

import (
	"newbeemall/controllers"
	"newbeemall/logic"
	"newbeemall/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddressAddHandler(c *gin.Context) {
	p := new(models.UserAddress)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("AddressAdd with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	userid, err := controllers.GetCurrentUser(c)
	if err != nil {
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	p.UserID = userid
	if err := logic.AddressAdd(p); err != nil {
		zap.L().Error("logic.AddressAdd failed", zap.Error(err))
		return
	}
	controllers.ResponseSuccess(c, "添加成功")
}
