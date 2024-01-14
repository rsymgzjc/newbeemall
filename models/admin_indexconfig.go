package models

type ParamIndexConfig struct {
	ConfigID    int64  `json:"config_id" db:"config_id"`
	ConfigName  string `json:"config_name" db:"configname"`
	ConfigType  int8   `json:"config_type" db:"configtype"`
	GoodsID     int64  `json:"goods_id" db:"goods_id"`
	RedirectURL string `json:"redirect_url" db:"redirecturl"`
	ConfigRank  int    `json:"config_rank" db:"configrank"`
	IsDeleted   int    `json:"isdeleted" db:"isdeleted"`
}

type ParamAddIndexConfig struct {
	ConfigName  string `json:"config_name" `
	ConfigType  int8   `json:"config_type" `
	GoodsID     int64  `json:"goods_id" `
	RedirectURL string `json:"redirect_url" `
	ConfigRank  int    `json:"config_rank" `
}
