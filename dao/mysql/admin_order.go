package mysql

import (
	"newbeemall/models"

	"github.com/jmoiron/sqlx"
)

func SearchOrder(p *models.UserIds) (datas []*models.ParamOrderStatus, err error) {
	datas = make([]*models.ParamOrderStatus, 0)
	sqlStr := `select order_id,orderstatus ,isdeleted from mallorder where order_id in (?)`
	query, args, err := sqlx.In(sqlStr, p.Ids)
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&datas, query, args...)
	return
}

func CheckDoneOrder(ids []int64) (err error) {
	sqlStr := `update mallorder set orderstatus =2 where order_id in (?)`
	query, args, err := sqlx.In(sqlStr, ids)
	if err != nil {
		return
	}
	query = db.Rebind(query)
	_, err = db.Exec(query, args...)
	return
}

func CheckOutOrder(ids []int64) (err error) {
	sqlStr := `update mallorder set orderstatus =3 where order_id in (?)`
	query, args, err := sqlx.In(sqlStr, ids)
	if err != nil {
		return
	}
	query = db.Rebind(query)
	_, err = db.Exec(query, args...)
	return
}

func CloseOrderOrder(ids []int64) (err error) {
	sqlStr := `update mallorder set orderstatus =-3 where order_id in (?)`
	query, args, err := sqlx.In(sqlStr, ids)
	if err != nil {
		return
	}
	query = db.Rebind(query)
	_, err = db.Exec(query, args...)
	return
}
