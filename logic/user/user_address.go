package user

import (
	"newbeemall/dao/mysql"
	"newbeemall/models"

	"go.uber.org/zap"
)

func AddressAdd(p *models.UserAddress) (err error) {
	if p.DefaultFlag == 1 {
		//判断用户已经存在的地址中是否有默认地址，如果有，就取消默认，否则直接添加
		if err = mysql.IsDefaultAddr(p); err != nil {
			zap.L().Error("设置失败", zap.Error(err))
			return
		}
	} else {
		//此地址不设为用户地址，直接添加即可
		if err = mysql.AddAddr(p); err != nil {
			zap.L().Error("直接添加失败", zap.Error(err))
			return
		}
	}
	return
}

func GetAddressList(userid int64) ([]*models.UserAddressList, error) {
	return mysql.GetAddressList(userid)
}

func UpdateAddr(p *models.UpdateAddr) (err error) {
	if p.DefaultFlag == 1 {
		if err = mysql.ToDefaultAddr(p.UserID); err != nil {
			return
		}
		if err = mysql.UpdateAddr(p); err != nil {
			return
		}
	} else {
		if err = mysql.UpdateAddr(p); err != nil {
			return
		}
	}
	return
}

func GetAddrDetail(id int64) (*models.UserAddrDetail, error) {
	return mysql.GetAddrDetail(id)
}

func GetDefAddr(userid int64) (*models.UserAddressList, error) {
	return mysql.GetDefAddr(userid)
}

func DelAddr(id int64, userid int64) (err error) {
	return mysql.DelAddr(id, userid)
}
