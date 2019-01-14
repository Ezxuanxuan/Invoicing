package models

import (
	"time"
)

type Components struct {
	Id        int       `xorm:"not null pk autoincr INT(11)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"index TIMESTAMP"`
	No        string    `xorm:"VARCHAR(30)"`
	Name      string    `xorm:"VARCHAR(30)"`
	Material  string    `xorm:"VARCHAR(30)"`
	Quality   int       `xorm:"INT(11)"` //质量
	Quantity  int       `xorm:"INT(11)"` //数量
}

//返回所有零件信息
func GetAllComponent() ([]Components, error) {
	components := make([]Components, 0)
	err := engine.Find(&components)
	if err != nil {
		return nil, err
	}
	return components, nil
}

//根据零件id返回零件信息
func GetComponentById(id int64) (Components, error) {
	component := new(Components)
	_, err := engine.Where("id = ?", id).Get(component)
	if err != nil {
		return Components{}, err
	}
	return *component, nil
}

//根据零件编号返回零件
func GetComponentByNo(no string) (bool, Components, error) {
	component := new(Components)
	has, err := engine.Where("no = ?", no).Get(component)
	if err != nil {
		return has, Components{}, err
	}
	return has, *component, nil
}

//添加零件
func CreateComponent(component Components) (int64, error) {
	affected, err := engine.InsertOne(&component)
	return affected, err
}

//是否存在该零件编号
func IsExsitComponentNo(no string) (bool, error) {
	component := new(Components)
	has, err := engine.Where("no = ?", no).Exist(component)
	return has, err
}

//通过id删除该零件信息
func DelComponentById(id int64) (int64, error) {
	component := new(Components)
	affected, err := engine.Where("id = ?", id).Delete(component)
	return affected, err
}

func IsExistComponentId(id int64) (bool, error) {
	component := new(Components)
	has, err := engine.Where("id = ?", id).Exist(component)
	return has, err
}
