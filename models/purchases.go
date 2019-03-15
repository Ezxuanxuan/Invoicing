package models

import (
	"time"
)

type Purchases struct {
	Id           int64     `xorm:"not null pk autoincr INT(11)"`
	CreatedAt    time.Time `xorm:"created"`
	UpdatedAt    time.Time `xorm:"updated"`
	DeletedAt    time.Time `xorm:"index TIMESTAMP"`
	Num          string    `xorm:"VARCHAR(255)"`
	ComponentId  int64     `xorm:"INT(11)"`
	Quantity     int64     `xorm:"INT(11)"`
	DeliveryDate time.Time `xorm:"TIMESTAMP"`
	OrderNo      string    `xorm:"VARCHAR(30)"`
}

func GetPurchaseById(id int64) (Purchases, bool, error) {
	purchase := new(Purchases)
	has, err := engine.Where("id = ?", id).Get(purchase)
	return *purchase, has, err
}

func GetPurchasesByOrder(orderId int64) ([]Purchases, bool, error) {
	purchases := make([]Purchases, 0)
	has, err := engine.Where("order_id = ?", orderId).Get(&purchases)
	return purchases, has, err
}

func ChangePurchase2Pro(id int64, count int64) (bool, bool, error) {
	session := engine.NewSession()
	err := session.Begin()
	defer session.Close()
	purchase := new(Purchases)
	has, err := session.Where("id = ?", id).Get(purchase)
	if err != nil {
		return false, false, err
	}
	if !has {
		return false, false, nil
	}

	if purchase.Quantity < count {
		return true, false, nil
	}
	purchase.Quantity = purchase.Quantity - count
	_, err = session.Update(purchase)

	if err != nil {
		session.Rollback()
		return true, true, err
	}

	_, err = InsertProductComponet(purchase.OrderNo, purchase.ComponentId, count, session)
	if err != nil {
		session.Rollback()
		return true, true, err
	}
	return true, true, nil

}

func InsertPuechaseComponet(order_no string, component_id int64, quantity int64) (int64, error) {
	purchase := new(Purchases)
	//查看该零件是否已存在
	has, err := engine.Where("order_no = ? and component_id = ?", order_no, component_id).Get(purchase)
	if err != nil {
		return 0, err
	}
	//如果该生产单中已经存在该零件id
	if has {
		purchase.Quantity = quantity + purchase.Quantity
		_, err := engine.Update(purchase)
		if err != nil {
			return 0, err
		}
		return 1, nil
	}
	purchase2 := new(Purchases)
	purchase2.OrderNo = order_no
	purchase2.ComponentId = component_id
	purchase2.Quantity = quantity
	return engine.InsertOne(purchase2)
}
