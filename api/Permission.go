package api

import (
	"github.com/Invoicing/error"
	"github.com/Invoicing/models"
	"github.com/labstack/echo"
	"strconv"
)

//创建用户权限
func CreatePermission() echo.HandlerFunc {
	return func(c echo.Context) error {
		StaffId := c.FormValue("staff_id")
		Context := c.FormValue("context")
		//转换用户id格式
		staff_id, err := strconv.ParseInt(StaffId, 10, 64)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}

		//检查该用户id是否存在
		has, err := models.IsExitStaffById(staff_id)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if !has {
			return sendError(errors.USER_NOT_EXI, c)
		}
		if len(Context) != 11 {
			return sendError(errors.POWER_INPUT_ERROR, c)
		}

		//检查该用户是否已有权限
		has, _ = models.IsExistPermissionByStaffId(staff_id)
		if has {
			return sendError(errors.POWER_EXIST_FOR_USER, c)
		}

		permission := models.Permissions{StaffId: staff_id, Context: Context}
		affected, err := models.CreatePermission(permission)
		if err != nil || affected < 1 {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, "", "创建权限成功", c)
	}
}

//修改用户权限
func UpdatePermission() echo.HandlerFunc {
	return func(c echo.Context) error {
		StaffId := c.FormValue("staff_id")
		Context := c.FormValue("context")
		//转换用户id格式
		staff_id, err := strconv.ParseInt(StaffId, 10, 64)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}

		//检查该id是否存在
		has, err := models.IsExitStaffById(staff_id)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if !has {
			return sendError(errors.USER_NOT_EXI, c)
		}

		//检查该用户是否已有权限
		has, _ = models.IsExistPermissionByStaffId(staff_id)
		if !has {
			return sendError(errors.POWER_NOT_EXIST_FOR_USER, c)
		}

		if len(Context) != 11 {
			return sendError(errors.POWER_INPUT_ERROR, c)
		}

		if len(Context) != 11 {
			return sendError(errors.POWER_INPUT_ERROR, c)
		}

		permission := models.Permissions{StaffId: staff_id, Context: Context}
		affected, err := models.UpdatePermission(staff_id, permission)
		if err != nil || affected < 1 {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, "", "更新权限成功", c)

	}
}

//获取员工权限通过员工id
func GetPermissionByStaffId() echo.HandlerFunc {
	return func(c echo.Context) error {
		StaffId := c.FormValue("staff_id")
		if StaffId == "" {
			return sendError(errors.USER_NOT_EXI, c)
		}
		staff_id, err := strconv.ParseInt(StaffId, 10, 64)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}

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
		Id := c.FormValue("id")
		id, err := strconv.ParseInt(Id, 10, 64)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}

		context, err := models.GetPermissionByStaff(id)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, context, "获取权限成功", c)
	}
}

//获取所有的权限信息
func GetAllPermission() echo.HandlerFunc {
	return func(c echo.Context) error {
		permissions, err := models.GetPermissions()
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, permissions, "以上为全部员工权限信息", c)
	}
}
