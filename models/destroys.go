package models

import (
	"time"
)

type Destroys struct {
	Id           int64     `xorm:"not null pk autoincr INT(11)" json:"id"`
	CreatedAt    time.Time `xorm:"TIMESTAMP"`
	UpdatedAt    time.Time `xorm:"TIMESTAMP"`
	DeletedAt    time.Time `xorm:"index TIMESTAMP"`
	ComponentId  int64     `xorm:"INT(11)"  json:"component_id"`
	Quantity     int64     `xorm:"INT(11)" json:"carries_quantity"`
	DeliveryDate time.Time `xorm:"TIMESTAMP"`
	OrderNo      string    `xorm:"VARCHAR(30)"`
}

type ComponentDes struct {
	Destroys   `xorm:"extends"`
	Components `xorm:"extends"`
}

func GetDesByOrder(order_no int64) ([]ComponentDes, error) {
	componentDes := make([]ComponentDes, 0)
	err := engine.Table("destroys").Join("INNER", "components",
		"components.id = destroys.component_id").Where("order_no = ?", order_no).Find(&componentDes)
	return componentDes, err
}

func ToInsertDesComponet(order_no string, component_id int64, quantity int64) (int64, error) {
	destroy := new(Destroys)
	//查看该零件是否已存在
	has, err := engine.Where("order_no = ? and component_id = ?", order_no, component_id).Get(destroy)
	if err != nil {
		return 0, err
	}
	//如果该生产单中已经存在该零件id
	if has {
		destroy.Quantity = quantity + destroy.Quantity
		_, err := engine.Update(destroy)
		if err != nil {
			return 0, err
		}
		return 1, nil
	}
	quality2 := new(Qualities)
	quality2.OrderNo = order_no
	quality2.ComponentId = component_id
	quality2.Quantity = quantity
	return engine.InsertOne(quality2)
}
