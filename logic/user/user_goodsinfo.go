package user

import (
	"newbeemall/dao/mysql"
	"newbeemall/models"

	"go.uber.org/zap"
)

func SearchGoodsInfo(page int64, categoryid int64, keyword string, orderby string) (datas []*models.ParamGoodsSearch, err error) {
	datas = make([]*models.ParamGoodsSearch, 0)
	datas1, err := mysql.SearchGoods(page, categoryid, keyword, orderby)
	if err != nil {
		zap.L().Error("查询商品信息失败", zap.Error(err))
		return
	}
	for _, data1 := range datas1 {
		data := &models.ParamGoodsSearch{
			GoodsID:       data1.GoodsID,
			GoodsName:     data1.GoodsName,
			GoodsIntro:    data1.GoodsIntro,
			GoodsCoverImg: data1.GoodsCoverImg,
			SellingPrice:  data1.SellingPrice,
		}
		datas = append(datas, data)
	}
	return datas, err
}

func GetGoodsDetail(goodid int64) (data *models.ParamGoodsInfoDetail, err error) {
	if err = mysql.GoodStatusIsOne(goodid); err != nil {
		zap.L().Error("商品已下架", zap.Error(err))
		return
	}
	data1, err := mysql.GetGoodDetail(goodid)
	if err != nil {
		zap.L().Error("商品信息出错", zap.Error(err))
		return
	}
	var list []string
	list = append(list, data1.GoodsCarousel)
	data = &models.ParamGoodsInfoDetail{
		GoodsID:           data1.GoodsID,
		GoodsName:         data1.GoodsName,
		GoodsIntro:        data1.GoodsIntro,
		GoodsCoverImg:     data1.GoodsCoverImg,
		SellingPrice:      data1.SellingPrice,
		GoodsDetail:       data1.GoodsDetail,
		OriginPrice:       data1.OriginPrice,
		Tag:               data1.Tag,
		GoodsCarouselList: list,
	}
	return data, err
}
