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

func JudgeStatus(ordernum int64) (err error) {
	var orderstatus int
	sqlStr := `select orderstatus from mallorder where isdeleted=0 and ordernum=?`
	err = db.Get(&orderstatus, sqlStr, ordernum)
	if orderstatus == 4 || orderstatus == -1 || orderstatus == -2 || orderstatus == -3 {
		return errors.New("订单状态异常")
	}
	return
}

func CancelOrder(order *models.ParamOrders) (err error) {
	sqlStr := `update mallorder set orderstatus=? ,update_time=? where ordernum=? and isdeleted=0`
	_, err = db.Exec(sqlStr, order.OrderStatus, order.UpdateTime, order.OrderNum)
	return
}

func GetOrderDetail(ordernum int64) (data *models.ParamOrders, err error) {
	data = new(models.ParamOrders)
	sqlStr := `select * from mallorder where ordernum=? and isdeleted=0`
	err = db.Get(data, sqlStr, ordernum)
	return
}

func GetOrderList(page, size, userid int64) (datas []*models.ParamOrders, err error) {
	datas = make([]*models.ParamOrders, 0)
	sqlStr := `select * from mallorder where user_id=? and isdeleted=0 order by update_time desc limit ?,?`
	err = db.Select(&datas, sqlStr, userid, (page-1)*size, size)
	return
}
