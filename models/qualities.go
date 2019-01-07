package models

import (
	"time"
)

type Qualities struct {
	Id          int       `xorm:"not null pk autoincr INT(11)"`
	CreatedAt   time.Time `xorm:"TIMESTAMP"`
	UpdatedAt   time.Time `xorm:"TIMESTAMP"`
	DeletedAt   time.Time `xorm:"index TIMESTAMP"`
	Num         string    `xorm:"VARCHAR(255)"`
	ComponentId int       `xorm:"INT(11)"`
	Quality     int       `xorm:"INT(11)"`
	Date        time.Time `xorm:"TIMESTAMP"`
}
