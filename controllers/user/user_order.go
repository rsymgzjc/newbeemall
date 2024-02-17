package user

import (
	"newbeemall/controllers"
	"newbeemall/logic/user"
	"newbeemall/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SaveOrderHandler(c *gin.Context) {
	p := new(models.ParamOrder)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SaveOrder with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	userid, err := controllers.GetCurrentUser(c)
	if err != nil {
		return
	}
	if err := user.SaveOrder(p, userid); err != nil {
		zap.L().Error("SaveOrder with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "生成订单成功")
}

func PaySuccessHandler(c *gin.Context) {
	ordernumstr := c.Query("ordernum")
	ordernum, err := strconv.ParseInt(ordernumstr, 10, 64)
	if err != nil {
		return
	}
	paytypestr := c.Query("paytype")
	paytype, err := strconv.Atoi(paytypestr)
	if err = user.PaySuccess(ordernum, paytype); err != nil {
		zap.L().Error("PaySuccess with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "支付订单成功")
}

func FinishOrderHandler(c *gin.Context) {
	ordernumstr := c.Param("ordernum")
	ordernum, err := strconv.ParseInt(ordernumstr, 10, 64)
	if err != nil {
		return
	}
	if err := user.FinishOrder(ordernum); err != nil {
		zap.L().Error("签收订单失败", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "签收订单成功")
}

func CancelOrderHandler(c *gin.Context) {
	ordernumstr := c.Param("ordernum")
	ordernum, err := strconv.ParseInt(ordernumstr, 10, 64)
	if err != nil {
		return
	}
	if err := user.CancelOrder(ordernum); err != nil {
		zap.L().Error("取消订单失败", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "取消订单失败")
}

func OrderDetailHandler(c *gin.Context) {
	ordernumstr := c.Param("ordernum")
	ordernum, err := strconv.ParseInt(ordernumstr, 10, 64)
	if err != nil {
		return
	}
	data, err := user.GetOrderDetail(ordernum)
	if err != nil {
		zap.L().Error("GetOrderDetail with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, data)
}

func GetOrderListHandler(c *gin.Context) {
	page, size := controllers.GetPageInfo(c)
	userid, err := controllers.GetCurrentUser(c)
	if err != nil {
		return
	}
	datas, err := user.GetOrderList(page, size, userid)
	if err != nil {
		zap.L().Error("GetOrderList with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	if len(datas) < 1 {
		controllers.ResponseSuccess(c, make([]interface{}, 0))
	} else {
		controllers.ResponseSuccess(c, datas)
	}
}
