package models

type ParamGoodsInfo struct {
	GoodsID         int64  `json:"goodsid" db:"goods_id"`
	GoodsName       string `json:"goodsname" db:"goodsname"`
	GoodsIntro      string `json:"goodsintro" db:"goodsintro"`
	GoodsCategoryID int64  `json:"goodscategoryid" db:"goodscategory_id"`
	GoodsCoverImg   string `json:"goodscoverimg" db:"goodscoverimg"`
	GoodsCarousel   string `json:"goodscarousel" db:"goodscarousel"`
	GoodsDetail     string `json:"goodsdetail" db:"goodsdetail"`
	OriginPrice     int    `json:"originprice" db:"originprice"`
	SellingPrice    int    `json:"sellingprice" db:"sellingprice"`
	StockNum        int    `json:"stocknum" db:"stocknum"`
	Tag             string `json:"tag" db:"tag"`
	GoodsSellStatus int    `json:"goodssellstatus" db:"goodssellstatus"`
}

type GoodsInfoList struct {
	Page      int64  `json:"page" form:"page"`
	Size      int64  `json:"size" form:"size"`
	GoodsName string `json:"goodsname" form:"goodsname" db:"goodsname"`
}
