package api

import (
	"github.com/Invoicing/error"
	"github.com/Invoicing/models"
	"github.com/labstack/echo"
	"strconv"
)

//v1.POST("/component/create", api.CreateComponent())
//v1.POST("/component/del", api.DelComponent())
//v1.POST("/component/all", api.GetAllComponent())

func CreateComponent() echo.HandlerFunc {
	return func(c echo.Context) error {
		No := c.FormValue("no")
		Name := c.FormValue("name")
		Material := c.FormValue("material")
		Quality := c.FormValue("quality")
		Quantity := c.FormValue("quantity")

		if No == "" {
			return sendError(errors.COMPONENT_NO_ERROR, c)
		}
		if Name == "" {
			return sendError(errors.COMPONENT_NAME_ERROR, c)
		}
		if Material == "" {
			return sendError(errors.COMPONENT_MATERIAL_ERROR, c)
		}
		//将质量转为int类型
		quality, err := strconv.ParseInt(Quality, 10, 64)
		if err != nil {
			return sendError(errors.COMPONENT_QUALITY_ERROR, c)
		}

		//将数量 转为int类型
		quantity, err := strconv.ParseInt(Quantity, 10, 64)
		if err != nil {
			return sendError(errors.COMPONENT_QUANTITY, c)
		}

		//查询零件编号是否存在
		isexsit, err := models.IsExsitComponentNo(No)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if isexsit == true {
			return sendError(errors.COMPONENT_NO_EXSIT, c)
		}

		component := models.Components{
			No:       No,
			Name:     Name,
			Material: Material,
			Quality:  quality,
			Quantity: quantity,
		}

		_, err = models.CreateComponent(component)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, "", "添加零件信息成功", c)
	}
}

func GetComponentById() echo.HandlerFunc {
	return func(c echo.Context) error {
		Id := c.FormValue("id")

		id, err := strconv.ParseInt(Id, 10, 64)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}

		has, err := models.IsExistComponentId(id)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if !has {
			return sendError(errors.COMPONENT_ID_NOT_EXIST, c)
		}

		component, err := models.GetComponentById(id)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, component, "获取零件信息成功", c)
	}
}

func GetComponentByNo() echo.HandlerFunc {
	return func(c echo.Context) error {
		No := c.FormValue("no")
		if No == "" {
			return sendError(errors.COMPONENT_NO_ERROR, c)
		}
		has, component, err := models.GetComponentByNo(No)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if !has {
			return sendError(errors.COMPONENT_NO_NOT_EXSIT, c)
		}
		return sendSuccess(1, component, "获取零件信息成功", c)
	}
}

func GetAllComponent() echo.HandlerFunc {
	return func(c echo.Context) error {
		components, err := models.GetAllComponent()
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, components, "以上为全部获取零件信息", c)
	}
}

func DelComponentById() echo.HandlerFunc {
	return func(c echo.Context) error {
		Id := c.FormValue("id")

		id, err := strconv.ParseInt(Id, 10, 64)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}

		has, err := models.IsExistComponentId(id)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if !has {
			return sendError(errors.COMPONENT_ID_NOT_EXIST, c)
		}

		affected, err := models.DelComponentById(id)
		if err != nil || affected < 1 {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, "", "删除零件信息成功", c)

	}
}
