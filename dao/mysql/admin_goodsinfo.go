package mysql

import (
	"newbeemall/models"

	"github.com/jmoiron/sqlx"
)

func GoodsInfoExist(name string, categoryid int64) (err error) {
	var count int
	sqlStr := `select count(goods_id) from goods_info where goodsname=? and goodscategory_id=?`
	err = db.Get(&count, sqlStr, name, categoryid)
	if count > 0 {
		return GoodsInfoExists
	}
	return
}

func CreateGoodsInfo(p *models.ParamGoodsInfo) (err error) {
	sqlStr := `insert into goods_info(goodsname, goodsintro, goodscategory_id, goodscoverimg, goodsdetail, originprice, sellingprice, stocknum, tag,goodssellstatus) 
				VALUES(?,?,?,?,?,?,?,?,?,?) `
	_, err = db.Exec(sqlStr, p.GoodsName, p.GoodsIntro, p.GoodsCategoryID, p.GoodsCoverImg, p.GoodsDetail, p.OriginPrice, p.SellingPrice, p.StockNum, p.Tag, p.GoodsSellStatus)
	return
}

func ChangeGoodsStatus(p *models.UserIds, status int64) (err error) {
	sqlStr := `update goods_info set goodssellstatus =? where goods_id in (?)`
	query, args, err := sqlx.In(sqlStr, status, p.Ids)
	if err != nil {
		return
	}
	query = db.Rebind(query)
	_, err = db.Exec(query, args...)
	return
}
