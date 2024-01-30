package user

import (
	"newbeemall/controllers"
	"newbeemall/logic/user"
	"newbeemall/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddShopCartHandler(c *gin.Context) {
	userid, err := controllers.GetCurrentUser(c)
	if err != nil {
		zap.L().Error("获取用户id失败", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeNeedLogin)
		return
	}
	p := new(models.ParamAddCart)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("AddShopCart with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	if err := user.AddShopCart(userid, p); err != nil {
		zap.L().Error("AddShopCart with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "添加成功")
}

func UpdateShopCartHandler(c *gin.Context) {
	p := new(models.ParamUpdateShopCart)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpdateShopCart with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	if err := user.UpdateShopCart(p); err != nil {
		zap.L().Error("UpdateShopCart with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "更新成功")
}

func DeleteShopCartHandler(c *gin.Context) {
	idStr := c.Param("cartid")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return
	}
	if err := user.DeleteShopCart(id); err != nil {
		zap.L().Error("DeleteShopCart with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, "删除成功")
}

func GetShopCartList(c *gin.Context) {
	userid, err := controllers.GetCurrentUser(c)
	if err != nil {
		return
	}
	datas, err := user.GetShopCartList(userid)
	if err != nil {
		zap.L().Error("GetShopCartList with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, datas)
}

func GetShopCartByIDs(c *gin.Context) {
	p := new(models.UserIds)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("GetShopCartByIDs with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	userid, err := controllers.GetCurrentUser(c)
	if err != nil {
		return
	}
	datas, err := user.GetShopCart(p, userid)
	if err != nil {
		zap.L().Error("GetShopCartList with some problems", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	controllers.ResponseSuccess(c, datas)
}
