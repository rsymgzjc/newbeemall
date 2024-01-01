package manager

import (
	"newbeemall/controllers"
	"newbeemall/logic/manager"
	"newbeemall/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateGoodsInfoHandler(c *gin.Context) {
	p := new(models.ParamGoodsInfo)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("CreateGoodsInfo with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	manager.CreateGoodsInfo(p)
}
