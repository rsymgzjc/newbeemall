package mysql

import (
	"errors"
	"newbeemall/models"
	"strings"
)

func SearchGoods(page int64, categoryid int64, keyword string, orderby string) (datas []*models.ParamGoodsInfo, err error) {
	datas = make([]*models.ParamGoodsInfo, 0)
	var sqlWhereParts []string
	var args []interface{}
	var sqlOrderPart string
	if keyword != "" {
		sqlWhereParts = append(sqlWhereParts, "(goodsname LIKE ? OR goodsintro LIKE ?)")
		args = append(args, "%"+keyword+"%", "%"+keyword+"%")
	}
	if categoryid >= 0 {
		sqlWhereParts = append(sqlWhereParts, "goodscategory_id = ?")
		args = append(args, categoryid)
	}
	Query := "select goods_id, goodsname, goodsintro, goodscategory_id, goodscoverimg, goodscarousel, goodsdetail, originprice, sellingprice, stocknum, tag, goodssellstatus from goods_info"
	if len(sqlWhereParts) > 0 {
		Query += " WHERE " + strings.Join(sqlWhereParts, " AND ")
	}
	switch orderby {
	case "new":
		sqlOrderPart = " ORDER BY goods_id DESC limit ?,?"
	case "price":
		sqlOrderPart = " ORDER BY sellingprice ASC limit ?,?"
	default:
		sqlOrderPart = " ORDER BY stocknum DESC limit ?,?"
	}
	Query += sqlOrderPart
	args = append(args, (page-1)*10, 10)
	err = db.Select(&datas, Query, args...)
	return
}

func GoodStatusIsOne(goodid int64) (err error) {
	var rmk int
	sqlStr := `select goodssellstatus from goods_info where goods_id=?`
	err = db.Get(&rmk, sqlStr, goodid)
	if rmk != 0 {
		return errors.New("商品已下架")
	}
	return
}

func GetGoodDetail(goodid int64) (data *models.ParamGoodsInfo, err error) {
	data = new(models.ParamGoodsInfo)
	sqlStr := `select goods_id, goodsname, goodsintro, goodscategory_id, goodscoverimg, goodscarousel, goodsdetail, originprice, sellingprice, stocknum, tag, goodssellstatus from goods_info where goods_id=?`
	err = db.Get(data, sqlStr, goodid)
	return
}
