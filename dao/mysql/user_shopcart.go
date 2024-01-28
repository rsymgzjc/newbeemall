package mysql

import "newbeemall/models"

func IsCartGoodsExists(userid int64, goodsid int64) (err error) {
	var count int
	sqlStr := `select count(cart_id) from shoppingcart where user_id=? and goods_id=? and isdeleted=0`
	err = db.Get(&count, sqlStr, userid, goodsid)
	if count > 0 {
		return CartGoodsExists
	}
	return
}

func ExceedGoodsTotal(userid int64) (err error) {
	var count int
	sqlStr := `select count(user_id) from shoppingcart where user_id=? and isdeleted=0`
	err = db.Get(&count, sqlStr, userid)
	if count > 20 {
		return ExceedCartTotal
	}
	return
}

func AddCartGoods(p *models.ParamShopCart) (err error) {
	sqlStr := `insert into shoppingcart(user_id, goods_id, goods_count) values (?,?,?)`
	_, err = db.Exec(sqlStr, p.UserID, p.GoodsID, p.GoodsCount)
	return
}
