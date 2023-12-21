package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"newbeemall/models"
)

const secret = "Orimiya123"

func UserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username=?`
	var count int
	err = db.Get(&count, sqlStr, username)
	if err != nil {
		return
	}
	if count > 0 {
		return UserExist1
	}
	return
}

func encrptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	h.Write([]byte(oPassword))
	return hex.EncodeToString(h.Sum(nil))
}
func UserInsert(user *models.User) (err error) {
	user.Password = encrptPassword(user.Password)
	sqlStr := `insert into user(user_id,password,username,email) values (?,?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Password, user.UserName, user.Email)
	return
}
