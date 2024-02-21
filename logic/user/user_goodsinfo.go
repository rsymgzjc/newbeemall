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
