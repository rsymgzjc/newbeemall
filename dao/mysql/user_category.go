package mysql

import (
	"newbeemall/models"

	"github.com/jmoiron/sqlx"
)

func SelectGoodsCategoryOne(levelone int) (datas []*models.AdminGoodsCategory, err error) {
	datas = make([]*models.AdminGoodsCategory, 0)
	sqlStr := `select category_id, categoryname, categorylevel, parentid, categoryrank, isdeleted from goods_category where parentid=0 and isdeleted=0 and categorylevel=?`
	err = db.Select(&datas, sqlStr, levelone)
	return
}

func SelectGoodsCategoryTwo(leveltwo int, firstlevel []int64) (datas []*models.AdminGoodsCategory, err error) {
	datas = make([]*models.AdminGoodsCategory, 0)
	sqlStr := `select category_id, categoryname, categorylevel, parentid, categoryrank, isdeleted from goods_category where parentid in (?) and isdeleted=0 and categorylevel=?`
	query, args, err := sqlx.In(sqlStr, firstlevel, leveltwo)
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&datas, query, args...)
	return
}

func SelectGoodsCategoryThree(levelthree int, secondlevel []int64) (datas []*models.AdminGoodsCategory, err error) {
	datas = make([]*models.AdminGoodsCategory, 0)
	sqlStr := `select category_id, categoryname, categorylevel, parentid, categoryrank, isdeleted from goods_category where parentid in (?) and isdeleted=0 and categorylevel=?`
	query, args, err := sqlx.In(sqlStr, secondlevel, levelthree)
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&datas, query, args...)
	return
}
