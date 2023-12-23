package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
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

func UserLogin(user *models.User) (err error) {
	opassword := encrptPassword(user.Password)
	sqlStr := `select user_id,username,password from user where username=?`
	err = db.Get(user, sqlStr, user.UserName)
	if errors.Is(err, sql.ErrNoRows) {
		return UserNotExist
	}
	if err != nil {
		return
	}
	if user.Password != opassword {
		return InvalidPassword
	}
	return
}
func UserUpdate(user *models.User, userid int64) (err error) {
	if user.Password != "" {
		password := encrptPassword(user.Password)
		sqlStr := `update user set username=?,password=?,introduction=?,gender=? where user_id=?`
		_, err = db.Exec(sqlStr, user.UserName, password, user.Introduction, user.Gender, userid)
	} else {
		sqlStr := `update user set username=?,introduction=?,gender=? where user_id=?`
		_, err = db.Exec(sqlStr, user.UserName, user.Introduction, user.Gender, userid)
	}
	return
}
