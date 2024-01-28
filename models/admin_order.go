package models

import "time"

type ParamOrders struct {
	OrderId     int       `json:"order_id" form:"orderid" db:"order_id"`
	OrderNum    string    `json:"ordernum" form:"ordernum" db:"ordernum"`
	UserId      int       `json:"user_id" form:"userid" db:"user_id"`
	TotalPrice  int       `json:"totalprice" form:"totalprice" db:"totalprice"`
	PayStatus   int       `json:"paystatus" form:"paystatus" db:"paystatus"` //支付状态:0.未支付,1.支付成功,-1:支付失败
	PayType     int       `json:"payType" form:"payType" db:"paytype"`       //0.无 1.支付宝支付 2.微信支付
	PayTime     time.Time `json:"pay_time" form:"paytime" db:"pay_time"`
	OrderStatus int       `json:"orderstatus" form:"orderstatus" db:"orderstatus"` //订单状态:0.待支付 1.已支付 2.配货完成 3:出库成功 4.交易成功 -1.手动关闭 -2.超时关闭 -3.商家关闭
	ExtraInfo   string    `json:"extrainfo" form:"extrainfo" db:"extrainfo"`
	IsDeleted   int       `json:"isdeleted" form:"isdeleted" db:"isdeleted"`
	CreateTime  time.Time `json:"createtime" form:"createtime" db:"create_time"`
	UpdateTime  time.Time `json:"updatetime" form:"updatetime" db:"update_time"`
}
