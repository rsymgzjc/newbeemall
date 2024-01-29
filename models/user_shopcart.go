package models

type ParamAddCart struct {
	GoodsCount int   `json:"goods_count" `
	GoodsID    int64 `json:"goods_id"`
}

type ParamShopCart struct {
	CartID     int64 `json:"cart_id" db:"cart_id"`
	UserID     int64 `json:"userid" db:"user_id"`
	GoodsCount int   `json:"goods_count" db:"goods_count"`
	GoodsID    int64 `json:"goods_id" db:"goods_id"`
}

type ParamUpdateShopCart struct {
	CartID     int64 `json:"cart_id" db:"cart_id"`
	GoodsCount int   `json:"goods_count" db:"goods_count"`
}
