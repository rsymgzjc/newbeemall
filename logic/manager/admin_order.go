package manager

import (
	"newbeemall/dao/mysql"
	"newbeemall/models"

	"go.uber.org/zap"
)

func CheckDoneOrder(p *models.UserIds) (err error) {
	datas, err := mysql.SearchOrder(p)
	if err != nil {
		return
	}
	var ids []int64
	if len(datas) > 0 {
		for _, data := range datas {
			if data.IsDeleted == 1 {
				zap.L().Error("订单已被删除")
			} else if data.OrderStatus != 1 {
				zap.L().Error("订单状态不是支付状态")
			} else {
				ids = append(ids, data.OrderId)
			}
		}
	}
	if err = mysql.CheckDoneOrder(ids); err != nil {
		zap.L().Error("发货失败", zap.Error(err))
		return
	}
	return
}
