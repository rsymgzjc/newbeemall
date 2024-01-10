package models

type ParamCarousel struct {
	CarouselID   int64  `json:"carousel_id" db:"carousel_id"`
	CarouselURL  string `json:"carousel_url" db:"carouselurl"`
	RedirectURL  string `json:"redirect_url" db:"redirecturl"`
	CarouselRank int    `json:"carousel_rank" db:"carouselrank"`
	IsDeleted    int    `json:"isdeleted" db:"isdeleted"`
}

type ParamUpdateCarousel struct {
	CarouselID   int64  `json:"carousel_id" db:"carousel_id"`
	CarouselURL  string `json:"carousel_url" db:"carouselurl"`
	RedirectURL  string `json:"redirect_url" db:"redirecturl"`
	CarouselRank int    `json:"carousel_rank" db:"carouselrank"`
}
