package models

type ParamSignUp struct {
	UserName   string `json:"username" binding:"required"`
	PassWord   string `json:"password" binding:"required"`
	RePassWord string `json:"repassword" binding:"required,eqfield=PassWord"`
	Email      string `json:"email" binding:"required,email"`
}
