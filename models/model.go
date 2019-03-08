package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

const (
	IN       = 1
	OUT      = 2
	PURCHASE = 3
	PRODUCT  = 4
	QUALITY  = 5
	DESTROY  = 6
	CARRY    = 7
)

var (
	engine *xorm.Engine
	err    error
)

func Init() {
	engine, err = xorm.NewEngine("mysql", "root:@(127.0.0.1:3306)/Invoicing?charset=utf8")
	if err != nil {
		fmt.Println(err)
		panic("连接数据库失败")
	}
}
