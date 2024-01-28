package mysql

func IsCartGoodsExists(userid int64, goodsid int64) (err error) {
	var count int
	sqlStr := `select count(cart_id) from shoppingcart where user_id=? and goods_id=?`
	err = db.Get(&count, sqlStr, userid, goodsid)
	if count > 0 {
		return CartGoodsExists
	}
	return
}
