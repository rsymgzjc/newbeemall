package mysql

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
