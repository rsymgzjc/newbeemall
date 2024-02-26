package user

import (
	"newbeemall/controllers"
	"newbeemall/logic/user"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func GetIndexInfoHandler(c *gin.Context) {
	dataCarouselIndex, err := user.GetCarouselIndex(5)
	if err != nil {
		zap.L().Error("GetCarouselIndex with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	dataGoodsHotIndex, err := user.GetConfigGoodsForIndex(3, 4)
	if err != nil {
		zap.L().Error("GetConfigGoodsForIndex with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	dataGoodsNewIndex, err := user.GetConfigGoodsForIndex(4, 5)
	if err != nil {
		zap.L().Error("GetConfigGoodsForIndex with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	dataGoodsRecommendIndex, err := user.GetConfigGoodsForIndex(5, 10)
	if err != nil {
		zap.L().Error("GetConfigGoodsForIndex with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	indexResults := map[string]interface{}{}
	indexResults["carousels"] = dataCarouselIndex
	indexResults["hotgoods"] = dataGoodsHotIndex
	indexResults["newgoods"] = dataGoodsNewIndex
	indexResults["recommendgoods"] = dataGoodsRecommendIndex
	controllers.ResponseSuccess(c, indexResults)
}
