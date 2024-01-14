package mysql

import "newbeemall/models"

func CreateIndexConfig(index *models.ParamIndexConfig) (err error) {
	sqlStr := `insert into indexconfig(configname, configtype, goods_id, redirecturl, configrank) VALUES (?,?,?,?,?)`
	_, err = db.Exec(sqlStr, index.ConfigName, index.ConfigType, index.GoodsID, index.RedirectURL, index.ConfigRank)
	return
}

func IndexConfigExists(configtype int8, goodsid int64) (err error) {
	var count int
	sqlStr := `select count(config_id) from indexconfig where configtype=? and goods_id=? and isdeleted=0`
	err = db.Get(&count, sqlStr, configtype, goodsid)
	if count > 0 {
		return IndexConfigExist
	}
	return
}
