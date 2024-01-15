package manager

import (
	"newbeemall/controllers"
	"newbeemall/logic/manager"
	"newbeemall/models"
	"strconv"

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

func DeleteIndexConfigHandler(c *gin.Context) {
	p := new(models.UserIds)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("DeleteIndexConfig with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	if err := manager.DeleteIndexConfig(p); err != nil {
		zap.L().Error("DeleteIndexConfig with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "删除成功")
}

func UpdateIndexConfigHandler(c *gin.Context) {
	p := new(models.ParamUpdateIndex)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpdateIndexConfig with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	if err := manager.UpdateIndexConfig(p); err != nil {
		zap.L().Error("UpdateIndexConfig with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "更新成功")
}

func GetIndexConfigByID(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return
	}
	data, err := manager.GetIndexConfigByID(id)
	if err != nil {
		zap.L().Error("GetIndexConfigByID with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, data)
}

func GetIndexConfigList(c *gin.Context) {
	page, size := controllers.GetPageInfo(c)
	datas, err := manager.GetIndexConfigList(page, size)
	if err != nil {
		zap.L().Error("GetIndexConfigList with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, datas)
}
