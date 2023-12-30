package user

import (
	"newbeemall/dao/mysql"
	"newbeemall/models"
	"newbeemall/pkg/jwt"
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
		return
	}
	return
}

func Login(p *models.ParamLogin) (Token string, err error) {
	if err = mysql.IsUserLock(p.Username); err != nil {
		zap.L().Error("用户被封了", zap.Error(err))
		return
	}
	user := &models.User{
		UserName: p.Username,
		Password: p.Password,
	}
	if err = mysql.UserLogin(user); err != nil {
		zap.L().Error("登录失败", zap.Error(err))
		return
	}
	return jwt.GenToken(user.UserID, user.UserName)
}

func Update(p *models.ParamUpdate, userid int64) (err error) {
	user := &models.User{
		UserName:     p.Username,
		Password:     p.Password,
		Introduction: p.Introduction,
		Gender:       p.Gender,
	}
	if err = mysql.UserExist(p.Username); err != nil {
		zap.L().Error("用户重名了", zap.Error(err))
		return
	}
	if err = mysql.UserUpdate(user, userid); err != nil {
		zap.L().Error("更新失败", zap.Error(err))
		return
	}
	return
}

func GetInfo(userid int64) (*models.ParamUserDetail, error) {
	return mysql.GetUserInfo(userid)
}
