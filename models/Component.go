package models

import "github.com/jinzhu/gorm"

//15.零件表
type Component struct {
	gorm.Model
	No       string `gorm:"type:varchar(20)"` //零件编号
	Name     string `gorm:"type:varchar(20)"` //零件名
	Material string `gorm:"type:varchar(20)"` //材质
	quality  int    //单品质量
}
