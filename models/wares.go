package models

import (
	"time"
)

type Wares struct {
	Id          int       `xorm:"not null pk autoincr INT(11)"`
	CreatedAt   time.Time `xorm:"TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"TIMESTAMP"`
	DeletedAt   time.Time `xorm:"index TIMESTAMP"`
	ComponentId int       `xorm:"INT(11)"`
	Quantity    int       `xorm:"INT(11)"`
}
