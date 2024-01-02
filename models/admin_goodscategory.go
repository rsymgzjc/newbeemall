package models

type AdminGoodsCategory struct {
	CategoryLevel int    `json:"categorylevel" db:"categorylevel"`
	ParentId      int    `json:"parentid" db:"parentid"`
	CategoryName  string `json:"categoryname" db:"categoryname"`
	CategoryRank  string `json:"categoryrank" db:"categoryrank"`
	IsDeleted     int    `json:"isdeleted" db:"isdeleted"`
}
