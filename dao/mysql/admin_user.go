package mysql

import "newbeemall/models"

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
