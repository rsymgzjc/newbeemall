package mysql

import (
	"database/sql"
	"errors"
	"newbeemall/models"

	"go.uber.org/zap"
)

func IsDefaultAddr(p *models.UserAddress) (err error) {
	var count int
	sqlStr := `select count(address_id) from user_address where user_id=? and defaultflag=1`
	err = db.Get(&count, sqlStr, p.UserID)
	if count > 0 { //里面有默认地址
		err = ToDefaultAddr(p.UserID)
		if err != nil {
			return
		}
		err = AddAddr(p)
		if err != nil {
			return
		}
	} else {
		err = AddAddr(p)
		if err != nil {
			return
		}
	}
	return
}

func AddAddr(p *models.UserAddress) (err error) {
	sqlStr := `insert into user_address(user_id, username, userphone, defaultflag,provincename, cityname, regionname, detailaddress)
             values (?,?,?,?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.UserID, p.UserName, p.UserPhone, p.DefaultFlag, p.ProvinceName, p.CityName, p.RegionName, p.DetailAddress)
	return
}

func ToDefaultAddr(userid int64) (err error) {
	sqlStr := `update user_address set defaultflag=0 where defaultflag=1 and user_id=?`
	_, err = db.Exec(sqlStr, userid)
	return
}

func GetAddressList(userid int64) (list []*models.UserAddressList, err error) {
	sqlStr := `select username, userphone,provincename, cityname, regionname, detailaddress from user_address where user_id=?`
	err = db.Select(&list, sqlStr, userid)
	if errors.Is(err, sql.ErrNoRows) {
		zap.L().Warn("未添加地址")
		err = nil
	}
	return
}
