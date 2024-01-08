package manager

import (
	"newbeemall/dao/mysql"
	"newbeemall/models"
)

func CreateCarousel(p *models.ParamCarousel) error {
	return mysql.CreateCarousel(p)
}
