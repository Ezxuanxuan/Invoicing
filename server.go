package main

import (
	"flag"
	"github.com/Invoicing/models"
	"github.com/Invoicing/routes"
)

var (
	// 命令行参数
	confPath = flag.String("conf", "./config/pro.yaml", "配置文件位置")
)

func main() {
	flag.Parse()
	models.Init(*confPath)
	route := routes.Init()
	route.Logger.Fatal(route.Start(":1236"))
}
