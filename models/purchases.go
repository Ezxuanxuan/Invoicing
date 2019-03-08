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
	Quality      int64     `xorm:"INT(11)"`
	DeliveryDate time.Time `xorm:"TIMESTAMP"`
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
