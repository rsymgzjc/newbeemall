package mysql

import (
	"database/sql"
	"errors"
	"newbeemall/models"
)

func CategoryExist(level int, name string) (err error) {
	var count int
	sqlStr := `select count(category_id) from goods_category where categorylevel=? and categoryname=? and isdeleted=0`
	err = db.Get(&count, sqlStr, level, name)
	if errors.Is(err, sql.ErrNoRows) {
		return GoodsCategoryExist1
	}
	return
}

func CreateCategory(p *models.AdminGoodsCategory) (err error) {
	sqlStr := `insert into goods_category(categoryname,categorylevel, parentid,categoryrank,isdeleted) values (?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.CategoryName, p.CategoryLevel, p.ParentId, p.CategoryRank, p.IsDeleted)
	return
}
