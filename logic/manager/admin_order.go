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

func CheckOutOrder(p *models.UserIds) (err error) {
	datas, err := mysql.SearchOrder(p)
	if err != nil {
		return
	}
	var ids []int64
	if len(datas) > 0 {
		for _, data := range datas {
			if data.IsDeleted == 1 {
				zap.L().Error("订单已被删除")
			} else if data.OrderStatus != 1 && data.OrderStatus != 2 {
				zap.L().Error("订单状态不是支付状态或者配货完成")
			} else {
				ids = append(ids, data.OrderId)
			}
		}
	}
	if err = mysql.CheckOutOrder(ids); err != nil {
		zap.L().Error("出库失败", zap.Error(err))
		return
	}
	return
}

func CloseOrder(p *models.UserIds) (err error) {
	datas, err := mysql.SearchOrder(p)
	if err != nil {
		return
	}
	var ids []int64
	if len(datas) > 0 {
		for _, data := range datas {
			if data.IsDeleted == 1 {
				zap.L().Error("订单已被删除")
			} else if data.OrderStatus == 3 || data.OrderStatus < 0 {
				zap.L().Error("订单不能进行关闭操作")
			} else {
				ids = append(ids, data.OrderId)
			}
		}
	}
	if err = mysql.CloseOrderOrder(ids); err != nil {
		zap.L().Error("关闭订单失败", zap.Error(err))
		return
	}
	return
}

func GetOrder(orderid int64) (data *models.OrderDetail, err error) {
	data1, err := mysql.GetOrderByID(orderid)
	if err != nil {
		zap.L().Error("获取订单信息失败", zap.Error(err))
		return
	}
	datas, err := mysql.GetOrderItem(orderid)
	if err != nil {
		zap.L().Error("获取订单项失败", zap.Error(err))
		return
	}
	if len(datas) > 0 {
		data = &models.OrderDetail{
			OrderId:     data1.OrderId,
			OrderNum:    data1.OrderNum,
			TotalPrice:  data1.TotalPrice,
			PayType:     data1.PayType,
			OrderStatus: data1.OrderStatus,
			CreateTime:  data1.CreateTime,
			OrderItem:   datas,
		}
	}
	return data, err
}

func GetOrderList(page int64, size int64, numstr string, statusstr string) (datas []*models.ParamOrders, err error) {
	return mysql.GetOrderList1(page, size, numstr, statusstr)
}
