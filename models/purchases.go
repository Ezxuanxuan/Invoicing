package models

import (
	"fmt"
	"time"
)

type Purchases struct {
	Id           int64     `xorm:"not null pk autoincr INT(11)" json:"id"`
	CreatedAt    time.Time `xorm:"created"`
	UpdatedAt    time.Time `xorm:"updated"`
	DeletedAt    time.Time `xorm:"index TIMESTAMP"`
	ComponentId  int64     `xorm:"INT(11)" json:"component_id"`
	Quantity     int64     `xorm:"INT(11)" json:"purchase_quantity"`
	DeliveryDate time.Time `xorm:"TIMESTAMP"`
	OrderNo      string    `xorm:"VARCHAR(30)"`
}

type ComponentPurchases struct {
	Purchases  `xorm:"extends"`
	Components `xorm:"extends"`
}

func GetPurchaseById(id int64) (ComponentPurchases, bool, error) {

	purchase := new(ComponentPurchases)
	has, err := engine.Table("purchases").Join("INNER", "components",
		"components.id = purchases.component_id").Where("purchases.id = ?", id).Get(purchase)
	return *purchase, has, err
}

func GetPurchasesByOrder(order_no int64) ([]ComponentPurchases, error) {
	purchases := make([]ComponentPurchases, 0)
	err := engine.Table("purchases").Join("INNER", "components",
		"components.id = purchases.component_id").Where("order_no = ?", order_no).Find(&purchases)
	fmt.Println(err)
	return purchases, err
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

	_, err = ToInsertProductComponet(purchase.OrderNo, purchase.ComponentId, count)
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

func UpdatePurchase(id int64, count int64) (bool, error) {
	purchase := new(Purchases)
	has, err := engine.Where("id = ?", id).Get(purchase)

	if err != nil {
		return false, err
	}
	if !has {
		return false, nil
	}
	if purchase.Quantity += count; purchase.Quantity < 0 {
		return false, nil
	}
	_, err = engine.Update(purchase)
	if err != nil {
		return true, err
	}
	return true, nil
}

func ChangePurchase2Out(id int64, count int64) (bool, bool, error) {
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

	_, err = ToInsertOutComponet(purchase.OrderNo, purchase.ComponentId, count)
	if err != nil {
		session.Rollback()
		return true, true, err
	}
	return true, true, nil
}
