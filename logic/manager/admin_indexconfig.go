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
