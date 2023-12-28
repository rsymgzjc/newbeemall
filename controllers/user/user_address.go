package user

import (
	"newbeemall/controllers"
	"newbeemall/logic/user"
	"newbeemall/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddressAddHandler(c *gin.Context) {
	p := new(models.UserAddress)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("AddressAdd with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	userid, err := controllers.GetCurrentUser(c)
	if err != nil {
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	p.UserID = userid
	if err := user.AddressAdd(p); err != nil {
		zap.L().Error("logic.AddressAdd failed", zap.Error(err))
		return
	}
	controllers.ResponseSuccess(c, "添加成功")
}

func GetAddressListHandler(c *gin.Context) {
	userid, err := controllers.GetCurrentUser(c)
	if err != nil {
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	list, err := user.GetAddressList(userid)
	if err != nil {
		zap.L().Error("logic.GetAddressList failed", zap.Error(err))
		return
	}
	controllers.ResponseSuccess(c, list)
}

func UpdateAddressHandler(c *gin.Context) {
	p := new(models.UpdateAddr)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("UpdateAddressAdd with invalid param", zap.Error(err))
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	userid, err := controllers.GetCurrentUser(c)
	if err != nil {
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	p.UserID = userid
	if err := user.UpdateAddr(p); err != nil {
		zap.L().Error("logic.UpdateAddr failed", zap.Error(err))
		return
	}
	controllers.ResponseSuccess(c, "更新地址成功")
}

func GetAddrDetailHandler(c *gin.Context) {
	//获取地址id
	idStr := c.Param("addressid")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	data, err := user.GetAddrDetail(id)
	if err != nil {
		zap.L().Error("logic.GetAddrDetail failed", zap.Error(err))
		return
	}
	controllers.ResponseSuccess(c, data)
}

func GetDefaultAddrHandler(c *gin.Context) {
	userid, err := controllers.GetCurrentUser(c)
	if err != nil {
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	data, err := user.GetDefAddr(userid)
	if err != nil {
		zap.L().Error("logic.GetDefAddr failed", zap.Error(err))
		return
	}
	controllers.ResponseSuccess(c, data)
}

func DelAddrHandler(c *gin.Context) {
	//获取地址id
	idStr := c.Param("addressid")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		controllers.ResponseError(c, controllers.CodeInvalidParam)
		return
	}
	userid, err := controllers.GetCurrentUser(c)
	if err != nil {
		controllers.ResponseError(c, controllers.CodeServerBusy)
		return
	}
	if err := user.DelAddr(id, userid); err != nil {
		zap.L().Error("logic.DelAddr failed", zap.Error(err))
		return
	}
	controllers.ResponseSuccess(c, "删除成功")
}
