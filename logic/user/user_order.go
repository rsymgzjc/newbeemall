package user

import (
	"newbeemall/dao/mysql"
	"newbeemall/models"
	"newbeemall/pkg/snowflake"
	"time"

	"go.uber.org/zap"
)

func SaveOrder(p *models.ParamOrder, userid int64) (err error) {
	ids := &models.UserIds{
		Ids: p.Ids,
	}
	PriceTotal := 0
	var CartIds []int64
	var Goodids []int64
	GoodsMap := make(map[int64]*models.ParamGoodsInfo)
	var GoodsInfo []*models.ParamGoodsInfo
	datas, err := GetShopCart(ids, userid)
	if err != nil {
		zap.L().Error("获取购物车信息失败", zap.Error(err))
		return
	}
	if len(datas) < 1 {
		zap.L().Error("没有数据", zap.Error(err))
		return
	} else {
		for _, data := range datas {
			PriceTotal = PriceTotal + data.GoodsCount*data.SellingPrice
			CartIds = append(CartIds, data.CartID)
			Goodids = append(Goodids, data.ParamGoodsInfo.GoodsID)
		}
		if PriceTotal < 1 {
			zap.L().Error("价格异常", zap.Error(err))
			return
		}
		_, err := GetDefAddr(userid)
		if err != nil {
			zap.L().Error("查询默认地址异常", zap.Error(err))
			return
		}
		//查询购物车对应的商品信息
		GoodsInfo, err = mysql.GetGoodsInfo(Goodids)
		if err != nil {
			zap.L().Error("商品查询出了问题", zap.Error(err))
			return
		}
		for _, good := range GoodsInfo {
			if good.GoodsSellStatus == 1 {
				zap.L().Error("商品已下架")
				return
			}
			GoodsMap[good.GoodsID] = good
		}
		for _, data := range datas {
			if _, ok := GoodsMap[data.ParamShopCart.GoodsID]; !ok {
				zap.L().Error("购物车数据异常")
				return
			}
			if data.GoodsCount > GoodsMap[data.ParamShopCart.GoodsID].StockNum {
				zap.L().Error("商品库存不足")
				return
			}
		}
		if len(CartIds) > 0 && len(Goodids) > 0 {
			if err := mysql.DeleteShoppingItem(CartIds, userid); err != nil {
				zap.L().Error("删除购物项失败", zap.Error(err))
				return
			}
			if err := mysql.UpdateStockNum(datas); err != nil {
				zap.L().Error("更新库存数量失败", zap.Error(err))
				return
			}
			ordernum := snowflake.GetID()
			Order := &models.ParamOrders{
				OrderNum:    ordernum,
				UserId:      userid,
				TotalPrice:  PriceTotal,
				PayStatus:   0,
				PayType:     0,
				OrderStatus: 0,
				ExtraInfo:   "",
			}
			if err := mysql.SaveOrder(Order); err != nil {
				zap.L().Error("生成订单失败", zap.Error(err))
				return
			}
		}
	}
	return
}

func PaySuccess(ordernum int64, paytype int) (err error) {
	if err = mysql.OrderStatusisone(ordernum); err != nil {
		zap.L().Error("订单状态异常", zap.Error(err))
		return
	}
	order := &models.ParamOrders{
		OrderNum:    ordernum,
		PayStatus:   1,
		PayType:     paytype,
		PayTime:     time.Now(),
		OrderStatus: 1,
	}
	if err = mysql.OrderPaySuccess(order); err != nil {
		zap.L().Error("支付订单失败", zap.Error(err))
		return
	}
	return
}
func FinishOrder(ordernum int64) (err error) {
	if err = mysql.IsOrderExists(ordernum); err != nil {
		zap.L().Error("未查询到记录", zap.Error(err))
		return
	}
	order := &models.ParamOrders{
		OrderNum:    ordernum,
		OrderStatus: 4,
		UpdateTime:  time.Now(),
	}
	if err = mysql.FinishOrder(order); err != nil {
		zap.L().Error("签收订单失败", zap.Error(err))
		return
	}
	return
}
