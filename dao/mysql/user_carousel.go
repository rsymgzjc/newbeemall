package mysql

import "newbeemall/models"

func GetCarouselForIndex(num int) (datas []*models.CarouselIndex, err error) {
	datas = make([]*models.CarouselIndex, 0)
	sqlStr := `select carouselurl,redirecturl from carousel where isdeleted=0 order by carouselrank desc limit ?`
	err = db.Select(&datas, sqlStr, num)
	return
}

func GetIndexConfig(configtype int, num int) (ids []int64, err error) {
	ids = make([]int64, 0)
	sqlStr := `select goods_id from indexconfig where configtype=? and isdeleted=0 order by configrank desc limit ?`
	err = db.Select(&ids, sqlStr, configtype, num)
	return
}
