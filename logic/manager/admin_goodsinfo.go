package manager

import (
	"errors"
	"newbeemall/dao/mysql"
	"newbeemall/models"

	"go.uber.org/zap"
)

func CreateGoodsInfo(p *models.ParamGoodsInfo) error {
	if err := mysql.CategoryNotExist(p.GoodsCategoryID); err != nil {
		if errors.Is(err, mysql.GoodsCategoryNotExist1) {
			zap.L().Error("分类信息异常", zap.Error(err))
			return err
		}
	}
	if err := mysql.GoodsInfoExist(p.GoodsName, p.GoodsCategoryID); err != nil {
		if errors.Is(err, mysql.GoodsInfoExists) {
			zap.L().Error("商品信息已存在", zap.Error(err))
			return err
		}
	}
	err := mysql.CreateGoodsInfo(p)
	if err != nil {
		zap.L().Error("创建商品信息失败", zap.Error(err))
		return err
	}
	return err
}

func ChangeGoodsStatus(p *models.UserIds, status int64) error {
	return mysql.ChangeGoodsStatus(p, status)
}
