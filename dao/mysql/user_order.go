package mysql

import (
	"errors"
	"newbeemall/models"
	"strconv"

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
	sqlStr := `insert into mallorder(order_id,ordernum, user_id, totalprice, paystatus, paytype,orderstatus, extrainfo) VALUES(?,?,?,?,?,?,?,?)`
	_, err = db.Exec(sqlStr, order.OrderId, order.OrderNum, order.UserId, order.TotalPrice, order.PayStatus, order.PayType, order.OrderStatus, order.ExtraInfo)
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
	sqlStr := `select order_id, ordernum, user_id, totalprice, paystatus, paytype, pay_time, orderstatus, extrainfo, isdeleted, create_time, update_time from mallorder where user_id=? and isdeleted=0 order by update_time desc limit ?,?`
	err = db.Select(&datas, sqlStr, userid, (page-1)*size, size)
	return
}

func SaveOrderItems(items []*models.ParamOrderItem) (err error) {
	for _, item := range items {
		sqlStr := `insert into orderitem(order_id, goods_id, goodsname, goodscoverimg, sellingprice, goods_count) VALUES (?,?,?,?,?,?)`
		_, err = db.Exec(sqlStr, item.OrderId, item.GoodsID, item.GoodsName, item.GoodsCoverImg, item.SellingPrice, item.GoodsCount)
	}
	return
}

func GetOrderByID(orderid int64) (data *models.ParamOrders, err error) {
	data = new(models.ParamOrders)
	sqlStr := `select order_id, ordernum, user_id, totalprice, paystatus, paytype, pay_time,orderstatus, extrainfo, isdeleted, create_time, update_time from mallorder where order_id=?`
	err = db.Get(data, sqlStr, orderid)
	return
}

func GetOrderItem(orderid int64) (datas []*models.ParamOrderItem, err error) {
	datas = make([]*models.ParamOrderItem, 0)
	sqlStr := `select order_id, goods_id, goodsname, goodscoverimg, sellingprice, goods_count from orderitem where order_id=?`
	err = db.Select(&datas, sqlStr, orderid)
	return
}

func GetOrderList1(page int64, size int64, numstr string, statusstr string) (datas []*models.ParamOrders, err error) {
	datas = make([]*models.ParamOrders, 0)
	if numstr != "" {
		ordernum, _ := strconv.ParseInt(numstr, 10, 64)
		if statusstr != "" {
			orderstatus, _ := strconv.ParseInt(statusstr, 10, 64)
			sqlStr := `select order_id, ordernum, user_id, totalprice, paystatus, paytype, pay_time,orderstatus, extrainfo, isdeleted, create_time, update_time from mallorder where ordernum=? and orderstatus=? limit ?,?`
			err = db.Select(&datas, sqlStr, ordernum, orderstatus, (page-1)*size, size)
			return
		} else {
			sqlStr := `select order_id, ordernum, user_id, totalprice, paystatus, paytype, pay_time,orderstatus, extrainfo, isdeleted, create_time, update_time from mallorder where ordernum=? limit ?,?`
			err = db.Select(&datas, sqlStr, ordernum, (page-1)*size, size)
			return
		}
	} else if numstr == "" {
		if statusstr != "" {
			orderstatus, _ := strconv.ParseInt(statusstr, 10, 64)
			sqlStr := `select order_id, ordernum, user_id, totalprice, paystatus, paytype, pay_time,orderstatus, extrainfo, isdeleted, create_time, update_time from mallorder where orderstatus=? limit ?,?`
			err = db.Select(&datas, sqlStr, orderstatus, (page-1)*size, size)
			return
		} else {
			sqlStr := `select order_id, ordernum, user_id, totalprice, paystatus, paytype, pay_time,orderstatus, extrainfo, isdeleted, create_time, update_time from mallorder limit ?,?`
			err = db.Select(&datas, sqlStr, (page-1)*size, size)
			return
		}
	}
	return
}
