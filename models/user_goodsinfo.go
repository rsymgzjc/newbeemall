package models

type ParamGoodsSearch struct {
	GoodsID       int64  `json:"goodsid" db:"goods_id"`
	GoodsName     string `json:"goodsname" db:"goodsname"`
	GoodsIntro    string `json:"goodsintro" db:"goodsintro"`
	GoodsCoverImg string `json:"goodscoverimg" db:"goodscoverimg"`
	SellingPrice  int    `json:"sellingprice" db:"sellingprice"`
}

type ParamGoodsInfoDetail struct {
	GoodsID           int64    `json:"goodsid" db:"goods_id"`
	GoodsName         string   `json:"goodsname" db:"goodsname"`
	GoodsIntro        string   `json:"goodsintro" db:"goodsintro"`
	GoodsCoverImg     string   `json:"goodscoverimg" db:"goodscoverimg"`
	SellingPrice      int      `json:"sellingprice" db:"sellingprice"`
	GoodsDetail       string   `json:"goodsdetail" db:"goodsdetail"`
	OriginPrice       int      `json:"originprice" db:"originprice"`
	Tag               string   `json:"tag" db:"tag"`
	GoodsCarouselList []string `json:"goodsCarouselList"`
}
