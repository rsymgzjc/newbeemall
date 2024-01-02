package manager

import (
	"go.uber.org/zap"
	"newbeemall/dao/mysql"
	"newbeemall/models"
)

func CreateCategory(p *models.AdminGoodsCategory) (err error) {
	if err = mysql.CategoryExist(p.CategoryLevel, p.CategoryName); err != nil {
		zap.L().Error("存在相同类别", zap.Error(err))
		return
	}
	err = mysql.CreateCategory(p)
	if err != nil {
		zap.L().Error("添加类别失败", zap.Error(err))
		return
	}
	return
}
