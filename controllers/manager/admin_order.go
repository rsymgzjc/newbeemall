package manager

import (
	"newbeemall/controllers"
	"newbeemall/logic/manager"
	"newbeemall/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CheckDoneOrderHandler(c *gin.Context) {
	p := new(models.UserIds)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("CheckDoneOrder with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	if err := manager.CheckDoneOrder(p); err != nil {
		zap.L().Error("CheckDoneOrder with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "发货成功")
}

func CheckOutOrderHandler(c *gin.Context) {
	p := new(models.UserIds)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("CheckOutOrder with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	if err := manager.CheckOutOrder(p); err != nil {
		zap.L().Error("CheckOutOrder with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "出库成功")
}

func CloseOrderHandler(c *gin.Context) {
	p := new(models.UserIds)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("CloseOrder with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	if err := manager.CloseOrder(p); err != nil {
		zap.L().Error("CloseOrder with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "关闭订单成功")
}

func GetOrderHandler(c *gin.Context) {
	orderidstr := c.Param("orderId")
	orderid, err := strconv.ParseInt(orderidstr, 10, 64)
	if err != nil {
		return
	}
	data, err := manager.GetOrder(orderid)
	if err != nil {
		zap.L().Error("GetOrder with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, data)
}

func GetOrderListHandler(c *gin.Context) {
	page, size := controllers.GetPageInfo(c)
	ordernumstr := c.Query("ordernum")
	orderstatusstr := c.Query("orderstatus")
	datas, err := manager.GetOrderList(page, size, ordernumstr, orderstatusstr)
	if err != nil {
		zap.L().Error("GetOrderList with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, datas)
}
