package models

import (
	"time"
)

type Logs struct {
	Id        int       `xorm:"not null pk autoincr INT(11)"`
	CreatedAt time.Time `xorm:"TIMESTAMP"`
	UpdatedAt time.Time `xorm:"TIMESTAMP"`
	DeletedAt time.Time `xorm:"index TIMESTAMP"`
	UserId    int       `xorm:"INT(11)"`
	Way       string    `xorm:"VARCHAR(255)"`
	Text      string    `xorm:"VARCHAR(255)"`
}
