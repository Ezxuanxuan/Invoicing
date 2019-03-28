package api

import (
	"github.com/Invoicing/error"
	"github.com/Invoicing/models"
	"github.com/labstack/echo"
	"strconv"
)

//获取某条采购信息，通过id
func GetPurchaseById() echo.HandlerFunc {
	return func(c echo.Context) error {
		Id := c.FormValue("id")

		if Id == "" {
			return sendError(errors.ID_ERROR, c)

		}
		id, err := strconv.ParseInt(Id, 10, 64)
		if err != nil {
			return sendError(errors.ID_ERROR, c)

		}
		purchase, has, err := models.GetPurchaseById(id)
		if err != nil {
			return sendError(errors.DO_ERROR, c)

		}
		if !has {
			return sendError(errors.PUCHASE_NOT_EXIST, c)
		}

		return sendSuccess(1, purchase, "以上为该条采购信息", c)

	}
}

//通过订单获取某条采购信息
func GetPurchasesByOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		orderId := c.FormValue("order_no")

		orderid, err := strconv.ParseInt(orderId, 10, 64)
		if err != nil {
			return sendError(errors.ID_ERROR, c)
		}

		purchases, err := models.GetPurchasesByOrder(orderid)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}

		return sendSuccess(1, purchases, "以上为该单号的所有采购信息", c)
	}
}

//创建采购单
func CreatePurchaseOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		orderNo := c.FormValue("no")
		orderTag := c.FormValue("tag")
		//查询输入是否非法
		if orderNo == "" || orderTag == "" {
			return sendError(errors.INPUT_ERROR, c)
		}

		//查询该订单编号是否存在
		has, err := models.IsExistOrderNo(orderNo, PURCHASE)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		} //已存在
		if has {
			return sendError(errors.ORDER_EXIST, c)
		}

		err = models.CreateAllOrder(orderNo, orderTag)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, "", "已经生成全部相关order", c)
	}
}

//修改某采购单状态，通过id
func ChangePurchase2ProById() echo.HandlerFunc {
	return func(c echo.Context) error {
		Id := c.FormValue("id")
		Count := c.FormValue("count")
		id, err := strconv.ParseInt(Id, 10, 64)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}
		count, err := strconv.ParseInt(Count, 10, 64)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}

		has, ok, err := models.ChangePurchase2Pro(id, count)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if !has {
			return sendError(errors.ID_NOT_EXIST, c)
		}
		if !ok {
			return sendError(errors.COUNT_BEYOND, c)
		}

		return sendSuccess(1, "", "更改状态成功", c)
	}
}

//向采购单插入零件
func InsertPurchase() echo.HandlerFunc {
	return func(c echo.Context) error {
		orderNo := c.FormValue("order_no")
		componentId := c.FormValue("component_id")
		Count := c.FormValue("count")

		//查询输入是否非法
		if orderNo == "" {
			return sendError(errors.INPUT_ERROR, c)
		}
		count, err := strconv.ParseInt(Count, 10, 64)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}
		component_id, err := strconv.ParseInt(componentId, 10, 64)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}

		_, err = models.InsertPuechaseComponet(orderNo, component_id, count)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}

		return sendSuccess(1, "", "添加零件成功", c)
	}
}

//修改采购单零件数量，通过采购单id
func UpdatePurchase() echo.HandlerFunc {
	return func(c echo.Context) error {
		Id := c.FormValue("id")
		Count := c.FormValue("count")
		id, err := strconv.ParseInt(Id, 10, 64)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}
		count, err := strconv.ParseInt(Count, 10, 64)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}

		has, err := models.UpdatePurchase(id, count)

		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if !has {
			return sendError(errors.COUNT_BEYOND, c)
		}

		return sendSuccess(1, "", "添加数量成功", c)
	}
}

//修改某采购单状态至出库单，通过id
func ChangePurchase2OutById() echo.HandlerFunc {
	return func(c echo.Context) error {
		Id := c.FormValue("id")
		Count := c.FormValue("count")
		id, err := strconv.ParseInt(Id, 10, 64)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}
		count, err := strconv.ParseInt(Count, 10, 64)
		if err != nil {
			return sendError(errors.INPUT_ERROR, c)
		}

		has, ok, err := models.ChangePurchase2Out(id, count)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if !has {
			return sendError(errors.ID_NOT_EXIST, c)
		}
		if !ok {
			return sendError(errors.COUNT_BEYOND, c)
		}

		return sendSuccess(1, "", "更改状态成功", c)
	}
}
