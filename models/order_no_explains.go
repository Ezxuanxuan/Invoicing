package models

type OrderNoExplains struct {
	OrderNo   string `xorm:"VARCHAR(30)"`
	OrderType int64  `xorm:"INT(11)"`
	Tag       string `xorm:"VARCHAR(255)"`
}

//创建某单，以及备注
func CreateOrder(No string, Type int64, Tag string) (int64, error) {
	order := new(OrderNoExplains)
	order.OrderNo = No
	order.OrderType = Type
	order.Tag = Tag
	return engine.Insert(&order)
}

//查询是否存在该单号
func IsExistOrderNo(no string) (bool, error) {
	order := new(OrderNoExplains)
	return engine.Where("order_no = ?", order).Exist(order)
}
