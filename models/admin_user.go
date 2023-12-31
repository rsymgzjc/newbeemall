package models

import "time"

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

type UserIds struct {
	Ids []int64 `json:"ids" form:"ids"`
}

type AdminDetail struct {
	AdminName  string    `json:"adminname" db:"adminname"`
	CreateTime time.Time `json:"createtime" db:"create_time"`
}
