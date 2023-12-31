package mysql

import (
	"database/sql"
	"errors"
	"newbeemall/models"

	"github.com/jmoiron/sqlx"
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
func CategoryNotExist(id int64) (err error) {
	var count int
	sqlStr := `select count(category_id) from goods_category where category_id=? and isdeleted=0`
	err = db.Get(&count, sqlStr, id)
	if count < 1 {
		return GoodsCategoryNotExist1
	}
	return
}
func CreateCategory(p *models.AdminGoodsCategory) (err error) {
	sqlStr := `insert into goods_category(categoryname,categorylevel, parentid,categoryrank,isdeleted) values (?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.CategoryName, p.CategoryLevel, p.ParentId, p.CategoryRank, p.IsDeleted)
	return
}

func UpdateCategory(p *models.AdminGoodsCategory) (err error) {
	sqlStr := `update goods_category set categoryname=? where category_id=? `
	_, err = db.Exec(sqlStr, p.CategoryName, p.CategoryID)
	return
}

func GetCategoryList(p *models.ParamSearchCategory) (data []*models.AdminGoodsCategory, err error) {
	sqlStr := `select category_id,categorylevel,parentid,categoryname,categoryrank,isdeleted from goods_category where categorylevel=? and parentid=? and isdeleted=0 limit ?,?`
	data = make([]*models.AdminGoodsCategory, 0)
	err = db.Select(&data, sqlStr, p.CategoryLevel, p.ParentId, (p.Page-1)*p.Size, p.Size)
	return
}

func GetCategory(id int64) (data *models.AdminGoodsCategory, err error) {
	data = new(models.AdminGoodsCategory)
	sqlStr := `select category_id,categorylevel,parentid,categoryname,categoryrank,isdeleted from goods_category where category_id=? and isdeleted=0 `
	err = db.Get(data, sqlStr, id)
	return
}

func DelCategory(p *models.UserIds) (err error) {
	sqlStr := `update goods_category set isdeleted=1 where category_id in (?)`
	query, args, err := sqlx.In(sqlStr, p.Ids)
	if err != nil {
		return
	}
	query = db.Rebind(query)
	_, err = db.Exec(query, args...)
	return
}
