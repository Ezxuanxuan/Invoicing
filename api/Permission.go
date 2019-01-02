package api

import (
	"github.com/Invoicing/cookie"
	"github.com/Invoicing/error"
	"github.com/Invoicing/models"
	"github.com/labstack/echo"
	"strconv"
)

//创建用户权限
func CreatePermission() echo.HandlerFunc {
	return func(c echo.Context) error {
		StaffId := c.Param("id")
		Context := c.Param("context")
		//转换用户id格式
		staff_id, err := strconv.Atoi(StaffId)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}

		if len(Context) != 11 {
			return sendError(errors.POWER_INPUT_ERROR, c)
		}

		permission := models.Permission{StaffId: staff_id, Context: Context}

		if !models.CreatePermission(permission) {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, "", "创建权限成功", c)
	}
}

//修改用户权限
func UpdatePermission() echo.HandlerFunc {
	return func(c echo.Context) error {
		StaffId := c.Param("id")
		Context := c.Param("context")
		//转换用户id格式
		staff_id, err := strconv.Atoi(StaffId)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}

		if len(Context) != 11 {
			return sendError(errors.POWER_INPUT_ERROR, c)
		}

		permission := models.Permission{StaffId: staff_id, Context: Context}

		if !models.UpdatePermission(permission) {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, "", "更新权限成功", c)

	}
}

//获取员工权限通过员工id
func GetPermissionByStaffId() echo.HandlerFunc {
	return func(c echo.Context) error {
		StaffId := c.Param("id")
		if StaffId == "" {
			return sendError(errors.USER_NOT_EXI, c)
		}
		staff_id := cookie.DecryptId(StaffId) // 解密后的id

		context, err := models.GetPermissionByStaff(staff_id)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, context, "获取权限成功", c)
	}
}

//获取员工权限，通过权限id
func GetPermissionById() echo.HandlerFunc {
	return func(c echo.Context) error {
		Id := c.Param("id")
		id, err := strconv.Atoi(Id)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}

		context, err := models.GetPermissionById(id)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, context, "获取权限成功", c)
	}
}

func GetAllPermission() echo.HandlerFunc {
	return func(c echo.Context) error {
		permissions, err := models.GetPermissions()
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, permissions, "以上为全部员工权限信息", c)
	}
}
