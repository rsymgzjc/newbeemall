package mysql

import (
	"newbeemall/models"

	"github.com/jmoiron/sqlx"
)

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

func CartGoodsNotExists(cartid int64) (err error) {
	var count int
	sqlStr := `select count(cart_id) from shoppingcart where cart_id=? and isdeleted=0`
	err = db.Get(&count, sqlStr, cartid)
	if count < 1 {
		return CartGoodsNotExist
	}
	return
}

func UpdateShopCart(p *models.ParamUpdateShopCart) (err error) {
	sqlStr := `update shoppingcart set goods_count=? where cart_id=? and isdeleted=0 `
	_, err = db.Exec(sqlStr, p.GoodsCount, p.CartID)
	return
}

func DeleteShopCart(cartid int64) (err error) {
	sqlStr := `update shoppingcart set isdeleted=1 where cart_id=? and isdeleted=0 `
	_, err = db.Exec(sqlStr, cartid)
	return
}

func GetCartGoods(userid int64) (datas []*models.ParamShopCart, err error) {
	datas = make([]*models.ParamShopCart, 0)
	sqlStr := `select cart_id,user_id,goods_id,goods_count from shoppingcart where user_id=? and isdeleted=0`
	err = db.Select(&datas, sqlStr, userid)
	return
}

func GetCartGoodsByids(p *models.UserIds, userid int64) (datas []*models.ParamShopCart, err error) {
	datas = make([]*models.ParamShopCart, 0)
	sqlStr := `select cart_id,user_id,goods_id,goods_count from shoppingcart where user_id =? and isdeleted=0 and cart_id in (?)`
	query, args, err := sqlx.In(sqlStr, userid, p.Ids)
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&datas, query, args...)
	return
}
