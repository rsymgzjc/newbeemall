package manager

import (
	"newbeemall/dao/mysql"
	"newbeemall/models"

	"go.uber.org/zap"
)

func CreateIndexConfig(p *models.ParamAddIndexConfig) (err error) {
	if err = mysql.GoodsInfoNotExist(p.GoodsID); err != nil {
		zap.L().Error("商品不存在", zap.Error(err))
		return
	}
	if err = mysql.IndexConfigExists(p.ConfigType, p.GoodsID); err != nil {
		zap.L().Error("存在首页配置项", zap.Error(err))
		return
	}
	index := &models.ParamIndexConfig{
		ConfigName:  p.ConfigName,
		ConfigType:  p.ConfigType,
		GoodsID:     p.GoodsID,
		RedirectURL: p.RedirectURL,
		ConfigRank:  p.ConfigRank,
	}
	if err = mysql.CreateIndexConfig(index); err != nil {
		zap.L().Error("创建失败", zap.Error(err))
		return
	}
	return
}

func DeleteIndexConfig(p *models.UserIds) (err error) {
	return mysql.DeleteIndexConfig(p)
}

func UpdateIndexConfig(p *models.ParamUpdateIndex) (err error) {
	if err = mysql.GoodsInfoNotExist(p.GoodsID); err != nil {
		zap.L().Error("商品不存在", zap.Error(err))
		return
	}
	if err = mysql.IndexConfigNotExists(p.ConfigID); err != nil {
		zap.L().Error("查询记录为空", zap.Error(err))
		return
	}
	indexconfig := &models.ParamIndexConfig{
		ConfigID:    p.ConfigID,
		ConfigName:  p.ConfigName,
		ConfigType:  p.ConfigType,
		GoodsID:     p.GoodsID,
		RedirectURL: p.RedirectURL,
		ConfigRank:  p.ConfigRank,
	}
	if err = mysql.UpdateIndexConfig(indexconfig); err != nil {
		zap.L().Error("更新失败", zap.Error(err))
		return
	}
	return
}

func GetIndexConfigByID(id int64) (*models.ParamIndexConfig, error) {
	return mysql.GetIndexConfigByID(id)
}

func GetIndexConfigList(page int64, size int64) ([]*models.ParamIndexConfig, error) {
	return mysql.GetIndexConfigList(page, size)
}
