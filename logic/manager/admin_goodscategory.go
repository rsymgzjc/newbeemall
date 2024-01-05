package manager

import (
	"newbeemall/dao/mysql"
	"newbeemall/models"

	"go.uber.org/zap"
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

func UpdateCategory(p *models.AdminGoodsCategory) (err error) {
	if err = mysql.CategoryExist(p.CategoryLevel, p.CategoryName); err != nil {
		zap.L().Error("存在相同类别", zap.Error(err))
		return
	}
	if err = mysql.UpdateCategory(p); err != nil {
		zap.L().Error("更新类别失败", zap.Error(err))
		return
	}
	return
}

func GetCategoryList(p *models.ParamSearchCategory) ([]*models.AdminGoodsCategory, error) {
	return mysql.GetCategoryList(p)
}

func GetCategory(id int64) (*models.AdminGoodsCategory, error) {
	return mysql.GetCategory(id)
}

func DelCategory(p *models.UserIds) error {
	return mysql.DelCategory(p)
}
