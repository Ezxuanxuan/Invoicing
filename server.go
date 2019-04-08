package main

import (
	"flag"
	"github.com/Invoicing/models"
	"github.com/Invoicing/routes"
)

var (
	// 命令行参数
	confPath = flag.String("conf", "", "配置文件位置")
)

func main() {
	flag.Parse()
	models.Init("/Users/xuan/mygo/src/github.com/Invoicing/config/test.yaml")
	route := routes.Init()
	route.Logger.Fatal(route.Start(":1236"))
}
