package manager

import (
	"newbeemall/controllers"
	"newbeemall/logic/manager"
	"newbeemall/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateIndexConfigHandler(c *gin.Context) {
	p := new(models.ParamAddIndexConfig)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("CreateIndexConfig with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	if err := manager.CreateIndexConfig(p); err != nil {
		zap.L().Error("manager.CreateIndexConfig with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "创建成功")
}
