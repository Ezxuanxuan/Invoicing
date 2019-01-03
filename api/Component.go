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
		No := c.Param("no")
		Name := c.Param("name")
		Material := c.Param("material")
		Quality := c.Param("quality")

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
		quality, err := strconv.Atoi(Quality)
		if err != nil {
			return sendError(errors.COMPONENT_QUALITY_ERROR, c)
		}

		//查询零件编号是否存在
		isexsit, err := models.IsExsitComponentNo(No)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if isexsit == true {
			return sendError(errors.COMPONENT_NO_EXSIT, c)
		}

		component := models.Component{
			No:       No,
			Name:     Name,
			Material: Material,
			Quality:  quality,
		}

		if !models.CreateComponent(component) {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, "", "添加零件信息成功", c)
	}
}

func GetComponentById() echo.HandlerFunc {
	return func(c echo.Context) error {
		Id := c.Param("id")

		id, err := strconv.Atoi(Id)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
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
		No := c.Param("no")
		if No == "" {
			return sendError(errors.COMPONENT_NO_ERROR, c)
		}
		component, err := models.GetComponentByNo(No)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
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
		Id := c.Param("id")

		id, err := strconv.Atoi(Id)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}

		if !models.DelComponentById(id) {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, "", "删除零件信息成功", c)

	}
}
