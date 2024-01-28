package user

import (
	"errors"
	"newbeemall/dao/mysql"
	"newbeemall/models"

	"go.uber.org/zap"
)

func AddShopCart(userid int64, p *models.ParamAddCart) (err error) {
	if p.GoodsCount < 1 {
		return errors.New("商品数量不能小于1")
	}
	if p.GoodsCount > 5 {
		return errors.New("超出商品的最大购买数量")
	}
	err = mysql.IsCartGoodsExists(userid, p.GoodsID)
	if errors.Is(err, mysql.CartGoodsExists) {
		zap.L().Error("无需重复添加", zap.Error(err))
		return
	}
	err = mysql.GoodsInfoNotExist(p.GoodsID)
	if errors.Is(err, mysql.GoodsInfoNotExists) {
		zap.L().Error("没有此商品", zap.Error(err))
		return
	}

}