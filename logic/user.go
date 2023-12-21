package logic

import (
	"newbeemall/dao/mysql"
	"newbeemall/models"
	"newbeemall/pkg/snowflake"

	"go.uber.org/zap"
)

func SignUp(p *models.ParamSignUp) (err error) {
	if err = mysql.UserExist(p.UserName); err != nil {
		zap.L().Error("用户已经存在了", zap.Error(err))
		return
	}
	userid := snowflake.GetID()
	user := &models.User{
		UserID:   userid,
		Password: p.PassWord,
		UserName: p.UserName,
		Email:    p.Email,
	}
	if err = mysql.UserInsert(user); err != nil {
		zap.L().Error("插入失败", zap.Error(err))
	}
	return
}
