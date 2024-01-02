package manager

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"newbeemall/controllers"
	"newbeemall/dao/mysql"
	"newbeemall/logic/manager"
	"newbeemall/models"
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
	}
	controllers.ResponseSuccess(c, "添加成功")
}

func UpdateCategoryHandler(c *gin.Context) {

}
