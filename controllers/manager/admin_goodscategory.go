package manager

import (
	"errors"
	"newbeemall/controllers"
	"newbeemall/dao/mysql"
	"newbeemall/logic/manager"
	"newbeemall/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateCategoryHandler(c *gin.Context) {
	p := new(models.AdminGoodsCategory)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("CreateCategory with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	if err := manager.CreateCategory(p); err != nil {
		zap.L().Error("manager.CreateCategory with invalid param", zap.Error(err))
		if errors.Is(err, mysql.GoodsCategoryExist1) {
			controllers.ResponseError(c, controllers.CodeGoodsCateGoryExist)
			return
		}
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "添加成功")
}

func UpdateCategoryHandler(c *gin.Context) {
	p := new(models.AdminGoodsCategory)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpdateCategory with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	if err := manager.UpdateCategory(p); err != nil {
		zap.L().Error("manager.UpdateCategory with invalid param", zap.Error(err))
		if errors.Is(err, mysql.GoodsCategoryExist1) {
			controllers.ResponseError(c, controllers.CodeGoodsCateGoryExist)
			return
		}
	}
	controllers.ResponseSuccess(c, "更新成功")
}

func GetCategoryListHandler(c *gin.Context) {
	p := new(models.ParamSearchCategory)
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("SearchCategory with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	datas, err := manager.GetCategoryList(p)
	if err != nil {
		zap.L().Error("manager.GetCategoryList with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, datas)
}
