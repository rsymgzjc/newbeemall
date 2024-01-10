package mysql

import (
	"newbeemall/models"

	"github.com/jmoiron/sqlx"
)

func CreateCarousel(p *models.ParamCarousel) (err error) {
	sqlStr := `insert into carousel(carouselurl, redirecturl, carouselrank) values (?,?,?)`
	_, err = db.Exec(sqlStr, p.CarouselURL, p.RedirectURL, p.CarouselRank)
	return
}

func DeleteCarousel(p *models.UserIds) (err error) {
	sqlStr := `update carousel set isdeleted=1 where carousel_id in (?)`
	query, args, err := sqlx.In(sqlStr, p.Ids)
	if err != nil {
		return
	}
	query = db.Rebind(query)
	_, err = db.Exec(query, args...)
	return
}

func UpdateCarousel(p *models.ParamUpdateCarousel) (err error) {
	sqlStr := `update carousel set carouselurl=? , redirecturl=?, carouselrank=? where carousel_id=?`
	_, err = db.Exec(sqlStr, p.CarouselURL, p.RedirectURL, p.CarouselRank, p.CarouselID)
	return
}

func IsCarouselExists(carouselid int64) (err error) {
	var count int
	sqlStr := `select count(carousel_id) from carousel where carousel_id=? and isdeleted =0`
	err = db.Get(&count, sqlStr, carouselid)
	if err != nil {
		return
	}
	if count < 1 {
		return CarouselNotExists
	}
	return
}

func GetCarouselByID(id int64) (data *models.ParamCarousel, err error) {
	data = new(models.ParamCarousel)
	sqlStr := `select carousel_id,carouselurl,redirecturl,carouselrank,isdeleted from carousel where carousel_id=?`
	err = db.Get(data, sqlStr, id)
	return
}

func GetCarouselList(page int64, size int64) (datas []*models.ParamCarousel, err error) {
	datas = make([]*models.ParamCarousel, 0)
	sqlStr := `select carousel_id,carouselurl,redirecturl,carouselrank,isdeleted from carousel limit ?,?`
	err = db.Select(&datas, sqlStr, (page-1)*size, size)
	return
}
