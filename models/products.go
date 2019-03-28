package models

import (
	"time"
)

type Products struct {
	Id           int64     `xorm:"not null pk autoincr INT(11)" json:"id"`
	CreatedAt    time.Time `xorm:"TIMESTAMP"`
	UpdatedAt    time.Time `xorm:"TIMESTAMP"`
	DeletedAt    time.Time `xorm:"index TIMESTAMP"`
	ComponentId  int64     `xorm:"INT(11)"  json:"component_id"`
	Quantity     int64     `xorm:"INT(11)"`
	DeliveryDate time.Time `xorm:"TIMESTAMP"`
	OrderNo      string    `xorm:"VARCHAR(30)"`
}

type ComponentPro struct {
	Products   `xorm:"extends"`
	Components `xorm:"extends"`
}

func GetProByOrder(order_no int64) ([]ComponentPro, error) {
	componentPro := make([]ComponentPro, 0)
	err := engine.Table("purchases").Join("INNER", "components",
		"components.id = products.component_id").Where("order_no = ?", order_no).Find(&componentPro)
	return componentPro, err
}

func InsertProductComponet(order_no string, component_id int64, quantity int64) (int64, error) {
	return ToInsertProductComponet(order_no, component_id, quantity)
}

func ToInsertProductComponet(order_no string, component_id int64, quantity int64) (int64, error) {
	product := new(Products)
	//查看该零件是否已存在
	has, err := engine.Where("order_no = ? and component_id = ?", order_no, component_id).Get(product)
	if err != nil {
		return 0, err
	}
	//如果该生产单中已经存在该零件id
	if has {
		product.Quantity = quantity + product.Quantity
		_, err := engine.Update(product)
		if err != nil {
			return 0, err
		}
		return 1, nil
	}
	product2 := new(Products)
	product2.OrderNo = order_no
	product2.ComponentId = component_id
	product2.Quantity = quantity
	return engine.InsertOne(product2)
}

func ChangePro2Qu(id int64, count int64) (bool, bool, error) {
	session := engine.NewSession()
	err := session.Begin()
	defer session.Close()
	pro := new(Products)
	has, err := session.Where("id = ?", id).Get(pro)
	if err != nil {
		return false, false, err
	}
	if !has {
		return false, false, nil
	}

	if pro.Quantity < count {
		return true, false, nil
	}
	pro.Quantity = pro.Quantity - count
	_, err = session.Update(pro)

	if err != nil {
		session.Rollback()
		return true, true, err
	}

	_, err = ToInsertQuComponet(pro.OrderNo, pro.ComponentId, count)
	if err != nil {
		session.Rollback()
		return true, true, err
	}
	return true, true, nil
}
