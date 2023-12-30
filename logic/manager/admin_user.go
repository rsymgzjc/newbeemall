package manager

import (
	"newbeemall/dao/mysql"
	"newbeemall/models"
	"newbeemall/pkg/jwt"
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

func AdminLogin(p *models.AdminLogin) (Token string, err error) {
	admin := &models.AdminUser{
		Password:  p.Password,
		AdminName: p.Adminname,
	}
	if err = mysql.AdminLogin(admin); err != nil {
		zap.L().Error("登录失败", zap.Error(err))
		return
	}
	return jwt.GenToken(admin.AdminID, admin.AdminName)
}

func AdminUpdateName(p *models.AdminUpdate, adminid int64) error {
	return mysql.AdminUpdateName(p, adminid)
}

func AdminUpdatePassword(p *models.AdminUpdate, adminid int64) error {
	return mysql.AdminUpdatePassword(p, adminid)
}

func GetUsersList(page int64, size int64) (data []*models.ParamUserDetail, err error) {
	return mysql.GetUsersList(page, size)
}

func LockUsers(p *models.UserIds, lock int64) (err error) {

}
