package main

import (
	"github.com/Invoicing/models"
	"github.com/Invoicing/routes"
)

func main() {
	models.Init()
	route := routes.Init()
	route.Logger.Fatal(route.Start(":1236"))
}
