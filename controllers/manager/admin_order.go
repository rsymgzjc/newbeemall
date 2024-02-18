package manager

import (
	"newbeemall/controllers"
	"newbeemall/logic/manager"
	"newbeemall/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CheckDoneOrderHandler(c *gin.Context) {
	p := new(models.UserIds)
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("CheckDoneOrder with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	if err := manager.CheckDoneOrder(p); err != nil {
		zap.L().Error("CheckDoneOrder with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "发货成功")
}
