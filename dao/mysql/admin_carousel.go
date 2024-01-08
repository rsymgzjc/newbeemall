package mysql

import "newbeemall/models"

func CreateCarousel(p *models.ParamCarousel) (err error) {
	sqlStr := `insert into carousel(carouselurl, redirecturl, carouselrank) values (?,?,?)`
	_, err = db.Exec(sqlStr, p.CarouselURL, p.RedirectURL, p.CarouselRank)
	return
}
