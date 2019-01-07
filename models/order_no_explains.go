package models

type OrderNoExplains struct {
	OrderNo   string `xorm:"VARCHAR(30)"`
	OrderType int    `xorm:"INT(11)"`
	Tag       string `xorm:"VARCHAR(255)"`
}
