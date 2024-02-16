package mysql

import (
	"errors"
	"newbeemall/models"

	"github.com/jmoiron/sqlx"
)

func DeleteShoppingItem(cartid []int64, userid int64) (err error) {
	sqlStr := `update shoppingcart set isdeleted =1 where cart_id in (?) and user_id=?`
	query, args, err := sqlx.In(sqlStr, cartid, userid)
	if err != nil {
		return
	}
	query = db.Rebind(query)
	_, err = db.Exec(query, args...)
	return
}

func UpdateStockNum(datas []*models.ShopCartDetail) (err error) {
	for _, data := range datas {
		sqlStr := `update goods_info set stocknum=? where goods_id =? and goodssellstatus=0  `
		_, err = db.Exec(sqlStr, data.StockNum-data.GoodsCount, data.ParamShopCart.GoodsID)
		return
	}
	return
}

func SaveOrder(order *models.ParamOrders) (err error) {
	sqlStr := `insert into mallorder (ordernum, user_id, totalprice, paystatus, paytype,orderstatus, extrainfo) VALUES(?,?,?,?,?,?,?)`
	_, err = db.Exec(sqlStr, order.OrderNum, order.UserId, order.TotalPrice, order.PayStatus, order.PayType, order.OrderStatus, order.ExtraInfo)
	return
}

func OrderStatusisone(ordernum int64) (err error) {
	sqlStr := `select orderstatus from mallorder where ordernum=?`
	var orderstatus int
	err = db.Get(&orderstatus, sqlStr, ordernum)
	if orderstatus != 0 {
		return errors.New("订单状态异常")
	}
	return
}

func OrderPaySuccess(order *models.ParamOrders) (err error) {
	sqlStr := `update mallorder set paytype=?,paystatus=?,pay_time=?,orderstatus=? where ordernum=?`
	_, err = db.Exec(sqlStr, order.PayType, order.PayStatus, order.PayTime, order.OrderStatus, order.OrderNum)
	return
}

func IsOrderExists(ordernum int64) (err error) {
	sqlStr := `select count(order_id) from mallorder where ordernum=? and isdeleted=0`
	var count int
	err = db.Get(&count, sqlStr, ordernum)
	if count < 1 {
		return errors.New("未查询到记录")
	}
	return
}

func FinishOrder(order *models.ParamOrders) (err error) {
	sqlStr := `update mallorder set orderstatus=? ,update_time=? where ordernum=? and isdeleted=0`
	_, err = db.Exec(sqlStr, order.OrderStatus, order.UpdateTime, order.OrderNum)
	return
}
