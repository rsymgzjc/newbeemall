package user

import (
	"newbeemall/controllers"
	"newbeemall/logic/user"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func GetGoodsCategoryHandler(c *gin.Context) {
	datas, err := user.GetGoodsCategory()
	if err != nil {
		zap.L().Error("GetGoodsCategory with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, datas)
}
