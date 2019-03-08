package models

import (
	"time"
)

type Logs struct {
	Id        int64     `xorm:"not null pk autoincr INT(11)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"index TIMESTAMP"`
	UserId    int64     `xorm:"INT(11)"`
	Way       string    `xorm:"VARCHAR(255)"`
	Text      string    `xorm:"VARCHAR(255)"`
}

//返回所有的日志
func GetLogs() (logs []*Logs, err error) {
	logs = make([]*Logs, 1)
	err = engine.Find(&logs)
	return
}

func GetLogByUser(userId int64) (log Logs) {

	return
}
