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
	err = mysql.ExceedGoodsTotal(userid)
	if errors.Is(err, mysql.ExceedCartTotal) {
		zap.L().Error("超出购物车总量", zap.Error(err))
		return
	}
	shopcart := &models.ParamShopCart{
		UserID:     userid,
		GoodsCount: p.GoodsCount,
		GoodsID:    p.GoodsID,
	}
	if err = mysql.AddCartGoods(shopcart); err != nil {
		zap.L().Error("添加购物车失败", zap.Error(err))
		return
	}
	return
}

func UpdateShopCart(p *models.ParamUpdateShopCart) (err error) {
	if p.GoodsCount > 5 {
		return errors.New("超出商品的最大购买数量")
	}
	err = mysql.CartGoodsNotExists(p.CartID)
	if errors.Is(err, mysql.CartGoodsNotExist) {
		zap.L().Error("商品不存在", zap.Error(err))
		return
	}
	if err = mysql.UpdateShopCart(p); err != nil {
		zap.L().Error("更新购物车失败", zap.Error(err))
		return
	}
	return
}

func DeleteShopCart(cartid int64) (err error) {
	err = mysql.CartGoodsNotExists(cartid)
	if errors.Is(err, mysql.CartGoodsNotExist) {
		zap.L().Error("商品不存在", zap.Error(err))
		return
	}
	if err := mysql.DeleteShopCart(cartid); err != nil {
		zap.L().Error("删除购物车商品失败", zap.Error(err))
		return
	}
	return
}

func GetShopCartList(userid int64) (datas []*models.ShopCartDetail, err error) {
	CartGoods, err := mysql.GetCartGoods(userid)
	if err != nil {
		return
	}
	datas = make([]*models.ShopCartDetail, 0)
	for _, CartGood := range CartGoods {
		good, err := mysql.GetGoodsInfoByID(CartGood.GoodsID)
		if err != nil {
			zap.L().Error("GetGoodsInfoByID failed", zap.Error(err))
			return
		}
		CartDetail := &models.ShopCartDetail{
			ParamGoodsInfo: good,
			ParamShopCart:  CartGood,
		}
		datas = append(datas, CartDetail)
	}
	return
}

func GetShopCart(p *models.UserIds, userid int64) (datas []*models.ShopCartDetail, err error) {
	CartGoods, err := mysql.GetCartGoodsByids(p, userid)
	if err != nil {
		return
	}
	datas = make([]*models.ShopCartDetail, 0)
	for _, CartGood := range CartGoods {
		good, err := mysql.GetGoodsInfoByID(CartGood.GoodsID)
		if err != nil {
			zap.L().Error("GetGoodsInfoByID failed", zap.Error(err))
			return
		}
		CartDetail := &models.ShopCartDetail{
			ParamGoodsInfo: good,
			ParamShopCart:  CartGood,
		}
		datas = append(datas, CartDetail)
	}
	return
}
