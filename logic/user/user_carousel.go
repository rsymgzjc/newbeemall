package user

import (
	"newbeemall/dao/mysql"
	"newbeemall/models"

	"go.uber.org/zap"
)

func GetCarouselIndex(num int) (datas []*models.CarouselIndex, err error) {
	datas, err = mysql.GetCarouselForIndex(num)
	if err != nil {
		zap.L().Error("获取轮播图失败", zap.Error(err))
		return
	}
	return datas, err
}

func GetConfigGoodsForIndex(configtype int, num int) (datas []*models.ConfigGoodsIndex, err error) {
	datas = make([]*models.ConfigGoodsIndex, 0)
	ids, err := mysql.GetIndexConfig(configtype, num)
	if err != nil {
		zap.L().Error("获取热销商品失败", zap.Error(err))
		return
	}
	goodsInfo, err := mysql.GetGoodsInfo(ids)
	if err != nil {
		zap.L().Error("获取商品信息失败", zap.Error(err))
		return
	}
	for _, good := range goodsInfo {
		GoodsIndex := &models.ConfigGoodsIndex{
			GoodsID:       good.GoodsID,
			GoodsName:     good.GoodsName,
			GoodsIntro:    good.GoodsIntro,
			GoodsCoverImg: good.GoodsCoverImg,
			SellingPrice:  good.SellingPrice,
			Tag:           good.Tag,
		}
		datas = append(datas, GoodsIndex)
	}
	return datas, err
}
