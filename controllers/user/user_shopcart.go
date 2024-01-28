package user

import (
	"newbeemall/controllers"
	"newbeemall/logic/user"
	"newbeemall/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddShopCartHandler(c *gin.Context) {
	userid, err := controllers.GetCurrentUser(c)
	if err != nil {
		zap.L().Error("获取用户id失败", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeNeedLogin)
		return
	}
	p := new(models.ParamAddCart)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("AddShopCart with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	if err := user.AddShopCart(userid, p); err != nil {
		zap.L().Error("AddShopCart with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "添加成功")
}
