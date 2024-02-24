package models

type IndexCategory struct {
	CategoryID          int64                 `json:"categoryid" db:"category_id"`
	CategoryLevel       int                   `json:"categorylevel" db:"categorylevel"`
	CategoryName        string                `json:"categoryname" db:"categoryname"`
	SecondLevelCategory []SecondLevelCategory `json:"secondLevelCategory"`
}

type SecondLevelCategory struct {
	CategoryID         int64                `json:"categoryid" db:"category_id"`
	ParentId           int                  `json:"parentid" db:"parentid"`
	CategoryLevel      int                  `json:"categorylevel" db:"categorylevel"`
	CategoryName       string               `json:"categoryname" db:"categoryname"`
	ThirdLevelCategory []ThirdLevelCategory `json:"thirdLevelCategory"`
}

type ThirdLevelCategory struct {
	CategoryID    int64  `json:"categoryid" db:"category_id"`
	CategoryLevel int    `json:"categorylevel" db:"categorylevel"`
	CategoryName  string `json:"categoryname" db:"categoryname"`
}
