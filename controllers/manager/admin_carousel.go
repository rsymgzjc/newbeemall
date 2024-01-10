package manager

import (
	"newbeemall/controllers"
	"newbeemall/logic/manager"
	"newbeemall/models"
	"strconv"

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

func DeleteCarouselHandler(c *gin.Context) {
	p := new(models.UserIds)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("DeleteCarousel with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	if err := manager.DeleteCarousel(p); err != nil {
		zap.L().Error("manager.DeleteCarousel with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "删除成功")
}

func UpdateCarouselHandler(c *gin.Context) {
	p := new(models.ParamUpdateCarousel)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpdateCarousel with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	if err := manager.UpdateCarousel(p); err != nil {
		zap.L().Error("UpdateCarousel with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "更新成功")
}

func GetCarouselByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return
	}
	data, err := manager.GetCarouselByID(id)
	if err != nil {
		zap.L().Error("manager.GetCarouselByID with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, data)
}

func GetCarouselsListHandler(c *gin.Context) {
	page, size := controllers.GetPageInfo(c)
	datas, err := manager.GetCarouselList(page, size)
	if err != nil {
		zap.L().Error("manager.GetCarouselList with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, datas)
}
