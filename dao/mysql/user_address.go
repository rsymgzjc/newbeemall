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

func UpdateAddr(p *models.UpdateAddr) (err error) {
	sqlStr := `update user_address set username=?,userphone=?,defaultflag=?,provincename=?,cityname=?,regionname=?,detailaddress=? where user_id=? and address_id=?`
	_, err = db.Exec(sqlStr, p.UserName, p.UserPhone, p.DefaultFlag, p.ProvinceName, p.CityName, p.RegionName, p.DetailAddress, p.UserID, p.AddressID)
	return
}

func GetAddrDetail(id int64) (data *models.UserAddrDetail, err error) {
	data = new(models.UserAddrDetail)
	sqlStr := `select username,userphone,defaultflag,provincename,cityname,regionname,detailaddress from user_address where address_id=?`
	err = db.Get(data, sqlStr, id)
	return
}

func GetDefAddr(userid int64) (data *models.UserAddressList, err error) {
	data = new(models.UserAddressList)
	sqlStr := `select username,userphone,provincename,cityname,regionname,detailaddress from user_address where defaultflag=1 and user_id=?`
	err = db.Get(data, sqlStr, userid)
	if errors.Is(err, sql.ErrNoRows) {
		zap.L().Warn("没有默认地址")
		err = nil
	}
	return
}
