package models

import (
	"time"
)

type Qualities struct {
	Id           int64     `xorm:"not null pk autoincr INT(11)" json:"id"`
	CreatedAt    time.Time `xorm:"TIMESTAMP"`
	UpdatedAt    time.Time `xorm:"TIMESTAMP"`
	DeletedAt    time.Time `xorm:"index TIMESTAMP"`
	ComponentId  int64     `xorm:"INT(11)"  json:"component_id"`
	Quantity     int64     `xorm:"INT(11)" json:"qualities_quantity"`
	DeliveryDate time.Time `xorm:"TIMESTAMP"`
	OrderNo      string    `xorm:"VARCHAR(30)"`
}

type ComponentQu struct {
	Qualities  `xorm:"extends"`
	Components `xorm:"extends"`
}

func GetQuByOrder(order_no int64) ([]ComponentQu, error) {
	componentQus := make([]ComponentQu, 0)
	err := engine.Table("qualities").Join("INNER", "components",
		"components.id = qualities.component_id").Where("order_no = ?", order_no).Find(&componentQus)
	return componentQus, err
}

func ToInsertQuComponet(order_no string, component_id int64, quantity int64) (int64, error) {
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

func ChangeQu2Car(id int64, count int64) (bool, bool, error) {
	session := engine.NewSession()
	err := session.Begin()
	defer session.Close()
	qu := new(Qualities)
	has, err := session.Where("id = ?", id).Get(qu)
	if err != nil {
		return false, false, err
	}
	if !has {
		return false, false, nil
	}

	if qu.Quantity < count {
		return true, false, nil
	}
	qu.Quantity = qu.Quantity - count
	_, err = session.Update(qu)

	if err != nil {
		session.Rollback()
		return true, true, err
	}

	_, err = ToInsertCarComponet(qu.OrderNo, qu.ComponentId, count)
	if err != nil {
		session.Rollback()
		return true, true, err
	}
	return true, true, nil
}

func ChangeQu2Des(id int64, count int64) (bool, bool, error) {
	session := engine.NewSession()
	err := session.Begin()
	defer session.Close()
	qu := new(Qualities)
	has, err := session.Where("id = ?", id).Get(qu)
	if err != nil {
		return false, false, err
	}
	if !has {
		return false, false, nil
	}

	if qu.Quantity < count {
		return true, false, nil
	}
	qu.Quantity = qu.Quantity - count
	_, err = session.Update(qu)

	if err != nil {
		session.Rollback()
		return true, true, err
	}

	_, err = ToInsertDesComponet(qu.OrderNo, qu.ComponentId, count)
	if err != nil {
		session.Rollback()
		return true, true, err
	}
	return true, true, nil
}
