package routes

import (
	"github.com/Invoicing/api"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {

	e := echo.New()

	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.POST("/staff/login", api.Login())
		v1.POST("/staff/create", api.CreateStaff())
		v1.POST("/staff/all", api.GetAllStaff())
		v1.POST("/staff/update/password", api.ModifyPassword())
		v1.POST("/staff/update/telephone", api.UpdateTelephone())

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

	}
	return e
}
