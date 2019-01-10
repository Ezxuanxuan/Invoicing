package api

import (
	"github.com/Invoicing/error"
	"github.com/Invoicing/models"
	"github.com/labstack/echo"
	"strconv"
	"strings"
)

func CreateInOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		No := c.FormValue("no")   //入库单编号
		Tag := c.FormValue("tag") //入库单备注
		var Type int64 = 1        //入库单类型

		//查询入库单是否存在
		if No == "" {
			return sendError(errors.Order_NOT_EXIST, c)
		}
		has, err := models.IsExistOrderNo(No)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if !has {
			return sendError(errors.Order_NOT_EXIST, c)
		}

		//备注必须有
		if Tag == "" {
			return sendError(errors.DO_ERROR, c)
		}
		affected, err := models.CreateOrder(No, Type, Tag)
		if err != nil || affected < 1 {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, "", "创建入库单成功", c)
	}
}

func InsertComponentIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		No := c.FormValue("no")                     //入库单编号
		Component_id := c.FormValue("component_id") //零件id
		Quantity := c.FormValue("quantity")         //零件数量
		var Status int64 = 0                        //未审核

		//查询入库单是否存在
		if No == "" {
			return sendError(errors.Order_NOT_EXIST, c)
		}
		has, err := models.IsExistOrderNo(No)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if !has {
			return sendError(errors.Order_NOT_EXIST, c)
		}

		//将零件id转换成int64
		component_id, err := strconv.ParseInt(Component_id, 10, 64)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}
		has, err = models.IsExistComponentId(component_id)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if !has {
			return sendError(errors.COMPONENT_ID_NOT_EXIST, c)
		}

		//将数量转换成int64
		quantity, err := strconv.ParseInt(Quantity, 10, 64)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}
		//插入
		affected, err := models.InsertComponet(No, component_id, quantity, Status)
		if err != nil || affected < 1 {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, "", "插入零件成功", c)
	}
}

func InsertComponentsIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		No := c.FormValue("no")                       //入库单编号
		Component_ids := c.FormValue("component_ids") //零件id数组
		Quantity := c.FormValue("quantity")           //零件数量
		var Status int64 = 0                          //未审核

		//将componentid的"，"分割的string转换成int数组的过程。
		if Component_ids == "" {
			return sendError(errors.INPUT_ERROR, c)
		}
		temp := strings.Split(Component_ids, ",")
		//var component_ids [len(temp)]int64
		component_ids := make([]int64, len(temp))
		var e error = nil
		for i := 0; i < len(temp); i++ {
			component_ids[i], e = strconv.ParseInt(temp[i], 10, 64)
			if e != nil {
				return sendError(errors.INPUT_ERROR, c)
			}
			//逐个判断零件id是否存在
			has, e := models.IsExistComponentId(component_ids[i])
			if e != nil {
				return sendError(errors.DO_ERROR, c)
			}
			if !has {
				return sendError(errors.COMPONENT_ID_NOT_EXIST, c)
			}
		}

		//查询入库单是否存在
		if No == "" {
			return sendError(errors.Order_NOT_EXIST, c)
		}
		has, err := models.IsExistOrderNo(No)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if !has {
			return sendError(errors.Order_NOT_EXIST, c)
		}

		//将零件id转换成int64
		quantity, err := strconv.ParseInt(Quantity, 10, 64)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}
		//插入
		affected, err := models.InsertComponents(No, component_ids, quantity, Status)
		if err != nil || affected < 1 {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, "", "插入零件成功", c)
	}
}

func VerbInById() echo.HandlerFunc {
	return func(c echo.Context) error {
		Id := c.FormValue("id")
		Status := c.FormValue("status")

		id, err := strconv.ParseInt(Id, 10, 64)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}
		var status int64 = 0
		if Status == "已通过" {
			status = 1
		} else if Status == "未通过" {
			status = -1
		}

		affected, err := models.UpdateInStatusById(id, status)
		if err != nil || affected < 1 {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, "", "更改审核状态成功", c)
	}
}

func VerbInByOrderNo() echo.HandlerFunc {
	return func(c echo.Context) error {
		OrderNo := c.FormValue("order_no")
		Status := c.FormValue("status")

		//查询入库单是否存在
		if OrderNo == "" {
			return sendError(errors.Order_NOT_EXIST, c)
		}
		has, err := models.IsExistOrderNo(OrderNo)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if !has {
			return sendError(errors.Order_NOT_EXIST, c)
		}

		var status int64 = 0
		if Status == "已通过" {
			status = 1
		} else if Status == "未通过" {
			status = -1
		}

		err = models.UpdateInStatusByOrderNo(OrderNo, status)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, "", "更改审核状态成功", c)
	}
}

func DelInById() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func UpdateInQuantityById() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func GetInByOrderNo() echo.HandlerFunc {
	return func(c echo.Context) error {
		OrderNo := c.FormValue("order_no")
		//查询入库单是否存在
		if OrderNo == "" {
			return sendError(errors.Order_NOT_EXIST, c)
		}
		has, err := models.IsExistOrderNo(OrderNo)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if !has {
			return sendError(errors.Order_NOT_EXIST, c)
		}

		ins, err := models.GetInByOrderNo(OrderNo)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, ins, "以上为该order中所有的零件", c)
	}
}

func GetInByOrderNoByStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		OrderNo := c.FormValue("order_no")
		Status := c.FormValue("status")

		var status int64 = 0
		if Status == "已通过" {
			status = 1
		} else if Status == "未通过" {
			status = -1
		}
		//查询入库单是否存在
		if OrderNo == "" {
			return sendError(errors.Order_NOT_EXIST, c)
		}
		has, err := models.IsExistOrderNo(OrderNo)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if !has {
			return sendError(errors.Order_NOT_EXIST, c)
		}

		ins, err := models.GetInByOrderNoByStatus(OrderNo, status)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, ins, "以上为该order中所有的零件", c)
	}
}