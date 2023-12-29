package models

type AdminLogin struct {
	Adminname string `json:"adminname" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type AdminSignup struct {
	Adminname string `json:"adminname" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
type AdminUser struct {
	AdminID   int64  `db:"admin_id"`
	Password  string `db:"password"`
	AdminName string `db:"adminname"`
}

type AdminUpdate struct {
	AdminName string `json:"adminname"`
	Password  string `json:"password"`
}
