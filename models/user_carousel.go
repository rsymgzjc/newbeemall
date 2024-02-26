package models

type CarouselIndex struct {
	CarouselURL string `json:"carousel_url" db:"carouselurl"`
	RedirectURL string `json:"redirect_url" db:"redirecturl"`
}

type ConfigGoodsIndex struct {
	GoodsID       int64  `json:"goodsid" db:"goods_id"`
	GoodsName     string `json:"goodsname" db:"goodsname"`
	GoodsIntro    string `json:"goodsintro" db:"goodsintro"`
	GoodsCoverImg string `json:"goodscoverimg" db:"goodscoverimg"`
	SellingPrice  int    `json:"sellingprice" db:"sellingprice"`
	Tag           string `json:"tag" db:"tag"`
}
