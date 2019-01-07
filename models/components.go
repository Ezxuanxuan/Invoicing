package models

import (
	"time"
)

type Components struct {
	Id        int       `xorm:"not null pk autoincr INT(11)"`
	CreatedAt time.Time `xorm:"TIMESTAMP"`
	UpdatedAt time.Time `xorm:"TIMESTAMP"`
	DeletedAt time.Time `xorm:"index TIMESTAMP"`
	No        string    `xorm:"VARCHAR(30)"`
	Name      string    `xorm:"VARCHAR(30)"`
	Material  string    `xorm:"VARCHAR(30)"`
	Quantity  int       `xorm:"INT(11)"`
}

//返回所有零件信息
func GetAllComponent() ([]Component, error) {
	var components []Component
	err := db.Select(&components).Error
	if err != nil {
		return nil, err
	}
	return components, nil
}

//根据零件id返回零件信息
func GetComponentById(id int) (Component, error) {
	var component Component
	err := db.Where("id = ?", id).Select(&component).Error
	if err != nil {
		return Component{}, err
	}
	return component, nil
}

//根据零件编号返回零件
func GetComponentByNo(No string) (Component, error) {
	var component Component
	err := db.Where("No = ?", No).Select(&component).Error
	if err != nil {
		return Component{}, err
	}
	return component, nil
}

//添加零件
func CreateComponent(component Component) bool {
	err := db.Create(&component).Error
	if err != nil {
		return false
	}
	return true
}

//是否存在该零件编号
func IsExsitComponentNo(no string) (bool, error) {
	count := 0
	err := db.Model(&Component{}).Where("no = ?", no).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

//通过id删除该零件信息
func DelComponentById(id int) bool {
	var component Component
	err := db.Where("id = ?", id).Delete(&component).Error
	if err != nil {
		return false
	}
	return true
}
