package manager

import (
	"errors"
	"newbeemall/dao/mysql"
	"newbeemall/models"

	"go.uber.org/zap"
)

func CreateCarousel(p *models.ParamCarousel) error {
	return mysql.CreateCarousel(p)
}

func DeleteCarousel(p *models.UserIds) error {
	return mysql.DeleteCarousel(p)
}

func UpdateCarousel(p *models.ParamUpdateCarousel) (err error) {
	if err = mysql.IsCarouselExists(p.CarouselID); err != nil {
		if errors.Is(err, mysql.CarouselNotExists) {
			zap.L().Error("carousel 不存在", zap.Error(err))
			return
		}
	}
	if err = mysql.UpdateCarousel(p); err != nil {
		zap.L().Error("更新失败", zap.Error(err))
		return
	}
	return
}

func GetCarouselByID(id int64) (*models.ParamCarousel, error) {
	return mysql.GetCarouselByID(id)
}

func GetCarouselList(page int64, size int64) ([]*models.ParamCarousel, error) {
	return mysql.GetCarouselList(page, size)
}
