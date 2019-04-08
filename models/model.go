package models

import (
	"fmt"
	"github.com/Invoicing/tools"
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

var conf *tools.Conf

var (
	engine *xorm.Engine
	err    error
)

func Init(confPath string) {
	conf = tools.GetConf(confPath)
	engine, err = xorm.NewEngine("mysql", fmt.Sprintf(("%s:%s@(%s)/%s?charset=utf8"),
		conf.User, conf.Password, conf.Host, conf.Dbname))
	if err != nil {
		fmt.Println(err)
		panic("连接数据库失败")
	}
}
