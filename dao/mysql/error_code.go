package mysql

import "errors"

var (
	UserExist1      = errors.New("用户已存在")
	UserNotExist    = errors.New("用户不存在")
	InvalidPassword = errors.New("无效密码")
	InvalidID       = errors.New("无效id")
	AdminExist1     = errors.New("管理员已存在")
	AdminNotExist   = errors.New("管理员不存在")
)
