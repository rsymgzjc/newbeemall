package manager

import (
	"newbeemall/controllers"
	"newbeemall/logic/manager"
	"newbeemall/models"
	"strconv"

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
	if err := manager.CreateGoodsInfo(p); err != nil {
		zap.L().Error("manager.CreateGoodsInfo with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "创建商品信息成功")
}

func ChangeGoodsStatusHandler(c *gin.Context) {
	str := c.Param("status")
	status, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return
	}
	p := new(models.UserIds)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("ChangeGoodsStatus with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	if err := manager.ChangeGoodsStatus(p, status); err != nil {
		zap.L().Error("manager.ChangeGoodsStatus with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "操作成功")
}

func UpdateGoodsInfoHandler(c *gin.Context) {
	p := new(models.ParamGoodsInfo)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpdateGoodsInfo with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	if err := manager.UpdateGoodsInfo(p); err != nil {
		zap.L().Error("manager.UpdateGoodsInfo with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "更新商品成功")
}

func GetGoodsInfoByIDHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return
	}
	data, err := manager.GetGoodInfoByID(id)
	if err != nil {
		zap.L().Error("manager.GetGoodInfoByID with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, data)
}

func GetGoodsListHandler(c *gin.Context) {
	status := c.Query("goodssellstatus")
	p := new(models.GoodsInfoList)
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetGoodsList with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	datas, err := manager.GetGoodsList(p, status)
	if err != nil {
		zap.L().Error("manager.GetGoodsList with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, datas)
}
