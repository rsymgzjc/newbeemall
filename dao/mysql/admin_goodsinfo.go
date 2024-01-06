package mysql

import (
	"newbeemall/models"
	"strconv"

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

func UpdateGoodsInfo(p *models.ParamGoodsInfo) (err error) {
	sqlStr := `update goods_info set goodsname=?,goodsintro=?, goodscategory_id=?,goodscoverimg=?,goodsdetail=?,originprice=?,sellingprice=?,stocknum=?,tag=?,goodssellstatus=? where goods_id=?`
	_, err = db.Exec(sqlStr, p.GoodsName, p.GoodsIntro, p.GoodsCategoryID, p.GoodsCoverImg, p.GoodsDetail, p.OriginPrice, p.SellingPrice, p.StockNum, p.Tag, p.GoodsSellStatus, p.GoodsID)
	return
}

func GetGoodsInfoByID(id int64) (data *models.ParamGoodsInfo, err error) {
	data = new(models.ParamGoodsInfo)
	sqlStr := `select goods_id,goodsname,goodsintro,goodscategory_id,goodscoverimg,goodsdetail,originprice,sellingprice,stocknum,tag,goodssellstatus from goods_info where goods_id=?`
	err = db.Get(data, sqlStr, id)
	if err != nil {
		return
	}
	return
}

func GetGoodsList(p *models.GoodsInfoList, status string) (datas []*models.ParamGoodsInfo, err error) {
	var sqlStr string
	if p.GoodsName != "" {
		if status != "" {
			statusint, _ := strconv.ParseInt(status, 10, 64)
			sqlStr = `select goods_id,goodsname,goodsintro,goodscategory_id,goodscoverimg,goodsdetail,originprice,sellingprice,stocknum,tag,goodssellstatus from goods_info 
						where goodsname=? and goodssellstatus=? limit ?,?`
			err = db.Select(&datas, sqlStr, p.GoodsName, statusint, (p.Page-1)*p.Size, p.Size)
			return
		} else {
			sqlStr = `select goods_id,goodsname,goodsintro,goodscategory_id,goodscoverimg,goodsdetail,originprice,sellingprice,stocknum,tag,goodssellstatus from goods_info 
						where goodsname=? limit ?,?`
			err = db.Select(&datas, sqlStr, p.GoodsName, (p.Page-1)*p.Size, p.Size)
			return
		}
	} else {
		if status != "" {
			statusint, _ := strconv.ParseInt(status, 10, 64)
			sqlStr = `select goods_id,goodsname,goodsintro,goodscategory_id,goodscoverimg,goodsdetail,originprice,sellingprice,stocknum,tag,goodssellstatus from goods_info 
						where goodssellstatus=? limit ?,?`
			err = db.Select(&datas, sqlStr, statusint, (p.Page-1)*p.Size, p.Size)
			return
		} else {
			sqlStr = `select goods_id,goodsname,goodsintro,goodscategory_id,goodscoverimg,goodsdetail,originprice,sellingprice,stocknum,tag,goodssellstatus from goods_info 
						limit ?,?`
			err = db.Select(&datas, sqlStr, (p.Page-1)*p.Size, p.Size)
			return
		}
	}
}
