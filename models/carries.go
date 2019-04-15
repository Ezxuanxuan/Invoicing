package models

import (
	"time"
)

type Carries struct {
	Id           int64     `xorm:"not null pk autoincr INT(11)" json:"id"`
	CreatedAt    time.Time `xorm:"TIMESTAMP"`
	UpdatedAt    time.Time `xorm:"TIMESTAMP"`
	DeletedAt    time.Time `xorm:"index TIMESTAMP"`
	ComponentId  int64     `xorm:"INT(11)"  json:"component_id"`
	Quantity     int64     `xorm:"INT(11)" json:"carries_quantity"`
	DeliveryDate time.Time `xorm:"TIMESTAMP"`
	OrderNo      string    `xorm:"VARCHAR(30)"`
}

type ComponentCar struct {
	Purchases `xorm:"extends"`
	Carries   `xorm:"extends"`
}

func GetCarByOrder(order_no int64) ([]ComponentDes, error) {
	componentDes := make([]ComponentDes, 0)
	err := engine.Table("purchases").Join("INNER", "components",
		"components.id = carries.component_id").Where("order_no = ?", order_no).Find(&componentDes)
	return componentDes, err
}

func ToInsertCarComponet(order_no string, component_id int64, quantity int64) (int64, error) {
	quality := new(Qualities)
	//查看该零件是否已存在
	has, err := engine.Where("order_no = ? and component_id = ?", order_no, component_id).Get(quality)
	if err != nil {
		return 0, err
	}
	//如果该生产单中已经存在该零件id
	if has {
		quality.Quantity = quantity + quality.Quantity
		_, err := engine.Update(quality)
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

func ChangeCar2Out(id int64, count int64) (bool, bool, error) {
	session := engine.NewSession()
	err := session.Begin()
	defer session.Close()
	car := new(Carries)
	has, err := session.Where("id = ?", id).Get(car)
	if err != nil {
		return false, false, err
	}
	if !has {
		return false, false, nil
	}

	if car.Quantity < count {
		return true, false, nil
	}
	car.Quantity = car.Quantity - count
	_, err = session.Update(car)

	if err != nil {
		session.Rollback()
		return true, true, err
	}

	_, err = ToInsertOutComponet(car.OrderNo, car.ComponentId, count)
	if err != nil {
		session.Rollback()
		return true, true, err
	}
	return true, true, nil
}

func ChangeCar2In(id int64, count int64) (bool, bool, error) {
	session := engine.NewSession()
	err := session.Begin()
	defer session.Close()
	car := new(Carries)
	has, err := session.Where("id = ?", id).Get(car)
	if err != nil {
		return false, false, err
	}
	if !has {
		return false, false, nil
	}

	if car.Quantity < count {
		return true, false, nil
	}
	car.Quantity = car.Quantity - count
	_, err = session.Update(car)

	if err != nil {
		session.Rollback()
		return true, true, err
	}

	_, err = ToInsertInComponet(car.OrderNo, car.ComponentId, count)
	if err != nil {
		session.Rollback()
		return true, true, err
	}
	return true, true, nil
}
