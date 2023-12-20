package logic

import (
	"newbeemall/dao/mysql"
	"newbeemall/models"

	"go.uber.org/zap"
)

func SignUp(p *models.ParamSignUp) (err error) {
	if err = mysql.UserExist(p.UserName); err != nil {
		zap.L().Error("用户已经存在了", zap.Error(err))
		return
	}
}
