package mysql

import (
	"database/sql"
	"errors"
	"newbeemall/models"

	"github.com/jmoiron/sqlx"
)

func AdminExist(adminname string) (err error) {
	sqlStr := `select count(admin_id) from admin_user where adminname=?`
	var count int
	err = db.Get(&count, sqlStr, adminname)
	if err != nil {
		return
	}
	if count > 0 {
		return AdminExist1
	}
	return
}

func AdminInsert(admin *models.AdminUser) (err error) {
	admin.Password = encrptPassword(admin.Password)
	sqlStr := `insert into admin_user(admin_id,password,adminname) values (?,?,?)`
	_, err = db.Exec(sqlStr, admin.AdminID, admin.Password, admin.AdminName)
	return
}

func AdminLogin(admin *models.AdminUser) (err error) {
	oPassword := encrptPassword(admin.Password)
	sqlStr := `select admin_id,adminname,password from admin_user where adminname=?`
	err = db.Get(admin, sqlStr, admin.AdminName)
	if errors.Is(err, sql.ErrNoRows) {
		return AdminNotExist
	}
	if err != nil {
		return
	}
	if oPassword != admin.Password {
		return InvalidPassword
	}
	return
}

func AdminUpdateName(p *models.AdminUpdate, adminid int64) (err error) {
	sqlStr := `update admin_user set adminname=? where admin_id=?`
	_, err = db.Exec(sqlStr, p.AdminName, adminid)
	return
}

func AdminUpdatePassword(p *models.AdminUpdate, adminid int64) (err error) {
	sqlStr := `update admin_user set password=? where admin_id=?`
	_, err = db.Exec(sqlStr, p.Password, adminid)
	return
}

func GetUsersList(page int64, size int64) (users []*models.ParamUserDetail, err error) {
	users = make([]*models.ParamUserDetail, 0)
	sqlStr := `select user_id,username,email,introduction,gender,lockflag from user order by create_time desc limit ?,?`
	err = db.Select(&users, sqlStr, (page-1)*size, size)
	return
}

func LockUsers(p *models.UserIds, locks int64) (err error) {
	sqlStr := `update user set lockflag=? where user_id in (?)`
	query, args, err := sqlx.In(sqlStr, locks, p.Ids)
	if err != nil {
		return err
	}
	query = db.Rebind(query)
	_, err = db.Exec(query, args...)
	return
}

func GetAdminDetail(adminid int64) (data *models.AdminDetail, err error) {
	data = new(models.AdminDetail)
	sqlStr := `select adminname,create_time from admin_user where admin_id=?`
	err = db.Get(data, sqlStr, adminid)
	return
}
