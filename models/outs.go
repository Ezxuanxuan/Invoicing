package models

import (
	"time"
)

type Outs struct {
	Id          int       `xorm:"not null pk autoincr INT(11)"`
	CreatedAt   time.Time `xorm:"TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"TIMESTAMP"`
	DeletedAt   time.Time `xorm:"index TIMESTAMP"`
	OrderNo     string    `xorm:"VARCHAR(30)"`
	ComponentId int       `xorm:"INT(11)"`
	Quantity    int       `xorm:"INT(11)"`
	Status      int       `xorm:"INT(11)"`
}
