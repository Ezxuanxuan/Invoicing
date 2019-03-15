package models

import (
	"github.com/go-xorm/xorm"
	"time"
)

type Products struct {
	Id          int64     `xorm:"not null pk autoincr INT(11)"`
	CreatedAt   time.Time `xorm:"TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"TIMESTAMP"`
	DeletedAt   time.Time `xorm:"index TIMESTAMP"`
	Num         string    `xorm:"VARCHAR(255)"`
	ComponentId int64     `xorm:"INT(11)"`
	Quantity    int64     `xorm:"INT(11)"`
	Date        time.Time `xorm:"TIMESTAMP"`
	OrderNo     string    `xorm:"VARCHAR(30)"`
}

func ToInsertProductComponet(order_no string, component_id int64, quantity int64) (int64, error) {
	session := engine.NewSession()
	err := session.Begin()
	defer session.Close()
	if err != nil {
		return 0, err
	}
	return InsertProductComponet(order_no, component_id, quantity, session)
}

func InsertProductComponet(order_no string, component_id int64, quantity int64, session *xorm.Session) (int64, error) {
	product := new(Products)
	//查看该零件是否已存在
	has, err := session.Where("order_no = ? and component_id = ?", order_no, component_id).Get(product)
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
	return session.InsertOne(product2)
}
