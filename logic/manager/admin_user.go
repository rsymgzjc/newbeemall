package manager

import (
	"newbeemall/dao/mysql"
	"newbeemall/models"
	"newbeemall/pkg/snowflake"

	"go.uber.org/zap"
)

func AdminSignup(p *models.AdminSignup) (err error) {
	if err = mysql.AdminExist(p.Adminname); err != nil {
		zap.L().Error("管理员已经存在了", zap.Error(err))
		return
	}
	adminid := snowflake.GetID()
	admin := &models.AdminUser{
		AdminID:   adminid,
		Password:  p.Password,
		AdminName: p.Adminname,
	}
	if err = mysql.AdminInsert(admin); err != nil {
		zap.L().Error("插入失败", zap.Error(err))
		return
	}
	return
}
