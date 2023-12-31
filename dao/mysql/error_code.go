package mysql

import "errors"

var (
	UserExist1             = errors.New("用户已存在")
	UserNotExist           = errors.New("用户不存在")
	UserLocked             = errors.New("用户被锁定")
	InvalidPassword        = errors.New("无效密码")
	InvalidID              = errors.New("无效id")
	AdminExist1            = errors.New("管理员已存在")
	AdminNotExist          = errors.New("管理员不存在")
	GoodsCategoryExist1    = errors.New("商品类别已存在")
	GoodsCategoryNotExist1 = errors.New("商品类别不存在")
	GoodsInfoExists        = errors.New("商品信息已存在")
)
