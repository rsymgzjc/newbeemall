package redis

import (
	"newbeemall/pkg/jwt"
)

func AddUserToken(token string) (err error) {
	pipeline := rdb.TxPipeline()
	pipeline.Set(token, 1, jwt.TokenExpireDuration)
	_, err = pipeline.Exec()
	return
}

func DeleteUserToken(token string) (err error) {
	pipeline := rdb.TxPipeline()
	pipeline.Del(token)
	_, err = pipeline.Exec()
	return
}

func SearchUserToken(token string) int64 {
	pipeline := rdb.TxPipeline()
	v := pipeline.Exists(token)
	_, _ = pipeline.Exec()
	return v.Val()
}
