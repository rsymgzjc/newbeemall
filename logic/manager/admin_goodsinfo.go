package manager

import "newbeemall/models"

func CreateGoodsInfo(p *models.ParamGoodsInfo) error {
	if err := mysql.GoodsInfoExist(p.GoodsName, p.GoodsCategoryID); err != nil {

	}
}