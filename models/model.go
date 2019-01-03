package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("mysql", "root:QQQQqqqq.1111@(127.0.0.1:3306)/Invoicing?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
}

//4.进货表
type In struct {
	gorm.Model
	OrderNo     string `gorm:"type:varchar(20)"` //订单编号
	ComponentId int    //Component外键
	Quantity    int    //数量
	Status      int    //审核状态
}

//6.出货表
type Out struct {
	gorm.Model
	OrderNo     string `gorm:"type:varchar(25)"` //订单编号
	ComponentId int    //Component外键
	Quantity    int    //数量
	Status      int    //审核状态
}

//16.单号对应备注表
type OrderNoExplain struct {
	OrderNo   string `gorm:"type:varchar(25)"` //订单编号
	OrderType int    //单号类型（0，1）
	Tag       string //备注
}

//7.仓库表
type Ware struct {
	gorm.Model
	ComponentId int
	Quantity    int
}

//15.采购表
type Process struct {
	gorm.Model
	Num          string    `gorm:"type:varchar(25)"` //过程编号
	ComponentId  int       //零件id
	Quality      int       //数量
	DeliveryDate time.Time //交货日期
	Status       int       //零件状态 '1:purchase,2:product,3:quality,4:destroy,5:carry'
}

//状态日志表
type StatusLog struct {
	gorm.Model
	Num         string `gorm:"type:varchar(25)"` //过程编号
	ComponentId int    //零件id
	Quality     int    //数量
	StaffId     int    //操作员工
	Way         string `gorm:"type:varchar(25)"` //操作方式"质检通过，销毁，采购，投产"
}

//13.销售表
type Sale struct {
}

//14.日志表
type Log struct {
	gorm.Model
	UserId int
	TaskId int
	Way    string `gorm:"tyoe:varchar(25)"`
	Text   string `gorm:"size:255"`
}
