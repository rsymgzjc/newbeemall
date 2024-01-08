package manager

import (
	"newbeemall/controllers"
	"newbeemall/logic/manager"
	"newbeemall/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateCarouselHandler(c *gin.Context) {
	p := new(models.ParamCarousel)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("CreateCarousel with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	if err := manager.CreateCarousel(p); err != nil {
		zap.L().Error("manager.CreateCarousel with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "创建成功")
}
