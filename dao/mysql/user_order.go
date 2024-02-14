package mysql

import (
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
