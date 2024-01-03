package models

type AdminGoodsCategory struct {
	CategoryID    int64  `json:"categoryid" db:"category_id"`
	CategoryLevel int    `json:"categorylevel" db:"categorylevel"`
	ParentId      int    `json:"parentid" db:"parentid"`
	CategoryName  string `json:"categoryname" db:"categoryname"`
	CategoryRank  int    `json:"categoryrank" db:"categoryrank"`
	IsDeleted     int    `json:"isdeleted" db:"isdeleted"`
}

type ParamSearchCategory struct {
	CategoryLevel int   `json:"categorylevel" form:"categorylevel" db:"categorylevel"`
	ParentId      int   `json:"parentid" form:"parentid" db:"parentid"`
	Page          int64 `json:"page" form:"page"`
	Size          int64 `json:"size" form:"size"`
}
