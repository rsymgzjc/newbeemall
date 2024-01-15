package mysql

import (
	"newbeemall/models"

	"github.com/jmoiron/sqlx"
)

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

func IndexConfigNotExists(configid int64) (err error) {
	var count int
	sqlStr := `select count(config_id) from indexconfig where config_id=?`
	err = db.Get(&count, sqlStr, configid)
	if count < 1 {
		return IndexConfigNotExist
	}
	return
}
func DeleteIndexConfig(p *models.UserIds) (err error) {
	sqlStr := `delete from indexconfig where config_id in (?)`
	query, args, err := sqlx.In(sqlStr, p.Ids)
	if err != nil {
		return
	}
	query = db.Rebind(query)
	_, err = db.Exec(query, args...)
	return
}

func UpdateIndexConfig(p *models.ParamIndexConfig) (err error) {
	sqlStr := `update indexconfig set configname=? ,configtype=?,goods_id=?,redirecturl=?,configrank=? where config_id=?`
	_, err = db.Exec(sqlStr, p.ConfigName, p.ConfigType, p.GoodsID, p.RedirectURL, p.ConfigRank, p.ConfigID)
	return
}

func GetIndexConfigByID(id int64) (data *models.ParamIndexConfig, err error) {
	data = new(models.ParamIndexConfig)
	sqlStr := `select config_id,configname,configtype,goods_id,redirecturl,configrank from indexconfig where config_id=?`
	err = db.Get(data, sqlStr, id)
	return
}

func GetIndexConfigList(page int64, size int64) (datas []*models.ParamIndexConfig, err error) {
	datas = make([]*models.ParamIndexConfig, 0)
	sqlStr := `select config_id,configname,configtype,goods_id,redirecturl,configrank from indexconfig limit ?,?`
	err = db.Select(&datas, sqlStr, (page-1)*size, size)
	return
}
