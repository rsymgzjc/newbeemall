package user

import (
	"newbeemall/controllers"
	"newbeemall/logic/user"
	"newbeemall/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SaveOrderHandler(c *gin.Context) {
	p := new(models.ParamOrder)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SaveOrder with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	userid, err := controllers.GetCurrentUser(c)
	if err != nil {
		return
	}
	if err := user.SaveOrder(p, userid); err != nil {
		zap.L().Error("SaveOrder with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "生成订单成功")
}
