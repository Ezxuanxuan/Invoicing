package models

import (
	"time"
)

type Logs struct {
	Id        int64     `xorm:"not null pk autoincr INT(11)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"index TIMESTAMP"`
	UserId    int64     `xorm:"INT(11)" json:"user_id"`
	Way       string    `xorm:"VARCHAR(255)"`
	Text      string    `xorm:"VARCHAR(255)"`
}

type StaffLogs struct {
	Logs   `xorm:"extends"`
	Staffs `xorm:"extends"`
}

//返回所有的日志
func GetLogs(pages int) ([]StaffLogs, error) {
	stafflogs := make([]StaffLogs, 0)
	err := engine.Table("logs").Join("INNER", "staffs",
		"staffs.id = logs.userId_id").Limit((pages-1)*10, pages*10).Find(&stafflogs)
	return stafflogs, err
}

//返回日志通过用户筛选
func GetLogByUser(userId int64, pages int) ([]StaffLogs, error) {
	stafflogs := make([]StaffLogs, 0)
	err := engine.Table("logs").Join("INNER", "staffs",
		"staffs.id = logs.userId_id").Where("logs.user_id = ?", userId).Limit((pages-1)*10, pages*10).Find(&stafflogs)
	return stafflogs, err
}

//返回日志通过日期筛选
func GetLogByDate(dateBegin time.Time, dateEnd time.Time, pages int) ([]StaffLogs, error) {
	stafflogs := make([]StaffLogs, 0)
	err := engine.Table("logs").Join("INNER", "staffs",
		"staffs.id = logs.userId_id").Where("created_at > ? and created_at < ?", dateBegin, dateEnd).Limit((pages-1)*10, pages*10).Find(&stafflogs)
	return stafflogs, err
}

//添加日志
func InsertLog(userId int64, way string, text string) error {
	log := new(Logs)
	log.UserId = userId
	log.Way = way
	log.Text = text
	_, err := engine.InsertOne(log)
	return err
}
