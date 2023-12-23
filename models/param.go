package models

type ParamSignUp struct {
	UserName   string `json:"username" binding:"required"`
	PassWord   string `json:"password" binding:"required"`
	RePassWord string `json:"repassword" binding:"required,eqfield=PassWord"`
	Email      string `json:"email" binding:"required,email"`
}

type User struct {
	UserID       int64  `db:"user_id"`
	Password     string `db:"password"`
	UserName     string `db:"username"`
	Email        string `db:"email"`
	Introduction string `db:"introduction"`
	Gender       int8   `db:"gender"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamUpdate struct {
}
