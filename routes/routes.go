package routes

import (
	"github.com/Invoicing/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() (e *echo.Echo) {

	e = echo.New()
	e.Use(middleware.CORS())
	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.POST("/staff/login", api.Login())
		v1.POST("/staff/create", api.CreateStaff())
		v1.POST("/staff/all", api.GetAllStaff())
		v1.POST("/staff/update/password", api.ModifyPassword())
		v1.POST("/staff/update/telephone", api.UpdateTelephone())
		v1.POST("/staff/del", api.DelStaffById())

		v1.POST("/permission/create", api.CreatePermission())
		v1.POST("/permission/update", api.UpdatePermission())
		v1.POST("/permission/getbyid", api.GetPermissionById())
		v1.POST("/permission/getall", api.GetAllPermission())
		v1.POST("/permission/getbystaff", api.GetPermissionByStaffId())

		v1.POST("/component/create", api.CreateComponent())
		v1.POST("/component/del", api.DelComponentById())
		v1.POST("/component/all", api.GetAllComponent())
		v1.POST("/component/get/id", api.GetComponentById())
		v1.POST("/component/get/no", api.GetComponentByNo())

		v1.POST("/in/create/order", api.CreateInOrder())
		v1.POST("/in/insert/one", api.InsertComponentIn())
		v1.POST("/in/insert/some", api.InsertComponentsIn())
		v1.POST("/in/verb/id", api.VerbInById())
		v1.POST("/in/verb/order", api.VerbInByOrderNo())
		v1.POST("/in/del/id", api.DelInById())
		v1.POST("/in/update/quantity/id", api.UpdateInQuantityById())
		v1.POST("/in/get/order", api.GetInByOrderNo())
		v1.POST("/in/get/order/status", api.GetInByOrderNoByStatus())

		v1.POST("/out/create/order", api.CreateOutOrder())
		v1.POST("/out/insert/one", api.InsertComponentOut())
		v1.POST("/out/insert/some", api.InsertComponentsOut())
		v1.POST("/out/verb/id", api.VerbOutById())
		v1.POST("/out/verb/order", api.VerbOutByOrderNo())
		v1.POST("/out/del/id", api.DelOutById())
		v1.POST("/out/update/quantity/id", api.UpdateOutQuantityById())
		v1.POST("/out/get/order", api.GetOutByOrderNo())
		v1.POST("/out/get/order/status", api.GetOutByOrderNoByStatus())

		v1.POST("/order/get/all", api.GetAllOrder())
		v1.POST("/order/get/type", api.GetOrderByType())

		v1.POST("/purchase/get/id", api.GetPurchaseById())
		v1.POST("/purchase/get/order", api.GetPurchasesByOrder())
		v1.POST("/purchase/create", api.CreatePurchaseOrder())
		v1.POST("/purchase/change2pro", api.ChangePurchase2ProById())
		v1.POST("/purchase/insert", api.InsertPurchase())
		v1.POST("/purchase/update", api.UpdatePurchase())
		v1.POST("/purchase/change2out", api.ChangePurchase2OutById())

		v1.POST("/product/get/order", api.GetProByOrder())
		v1.POST("/product/create", api.CreateProOrder())
		v1.POST("/product/pro2qu", api.ChangePro2Qu())
		v1.POST("/product/insert", api.InsertPro())

		v1.POST("/quality/get/order", api.GetQuByOrder())
		v1.POST("/quality/qu2car", api.ChangeQu2Car())
		v1.POST("/quality/qu2des", api.ChangeQu2Des())

		v1.POST("/destroy/get/order", api.GetDesByOrder())

		v1.POST("/carry/get/order", api.GetCarByOrder())
		v1.POST("/carry/qu2car", api.ChangeCar2In())
		v1.POST("/carry/qu2des", api.ChangeCar2Out())
	}
	return
}
