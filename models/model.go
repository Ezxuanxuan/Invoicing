package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	engine *xorm.Engine
	err    error
)

func Init() {
	engine, err = xorm.NewEngine("mysql", "root:QQQQqqqq.1111@/Invoicing?charset=utf8")
	if err != nil {
		panic("连接数据库失败")
	}
}
