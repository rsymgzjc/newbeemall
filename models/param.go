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
	Username     string `json:"username"`
	Password     string `json:"password"`
	Introduction string `json:"introduction"`
	Gender       int8   `json:"gender"`
}

type ParamUserDetail struct {
	UserName     string `json:"username" db:"username"`
	Email        string `json:"email" db:"email"`
	Introduction string `json:"introduction" db:"introduction"`
	Gender       int8   `json:"gender" db:"gender"`
}

type UserAddress struct {
	UserID        int64  `json:"userid" db:"user_id"`
	UserName      string `json:"username" binding:"required" db:"username"`
	UserPhone     string `json:"userphone" binding:"required" db:"userphone"`
	DefaultFlag   int8   `json:"defaultflag"  db:"defaultflag"`
	ProvinceName  string `json:"provincename" binding:"required" db:"provincename"`
	CityName      string `json:"cityname" binding:"required" db:"cityname"`
	RegionName    string `json:"regionname" binding:"required" db:"regionname"`
	DetailAddress string `json:"detailaddress" binding:"required" db:"detailaddress"`
}

type UserAddressList struct {
	UserName      string `json:"username" db:"username"`
	UserPhone     string `json:"userphone"  db:"userphone"`
	ProvinceName  string `json:"provincename" binding:"required" db:"provincename"`
	CityName      string `json:"cityname" binding:"required" db:"cityname"`
	RegionName    string `json:"regionname" binding:"required" db:"regionname"`
	DetailAddress string `json:"detailaddress" binding:"required" db:"detailaddress"`
}

type UpdateAddr struct {
	AddressID     int    `json:"addressid" binding:"required" db:"address_id"`
	UserID        int64  `json:"userid" binding:"required" db:"user_id"`
	UserName      string `json:"username" binding:"required" db:"username"`
	UserPhone     string `json:"userphone" binding:"required" db:"userphone"`
	DefaultFlag   int8   `json:"defaultflag"  db:"defaultflag"`
	ProvinceName  string `json:"provincename" binding:"required" db:"provincename"`
	CityName      string `json:"cityname" binding:"required" db:"cityname"`
	RegionName    string `json:"regionname" binding:"required" db:"regionname"`
	DetailAddress string `json:"detailaddress" binding:"required" db:"detailaddress"`
}
