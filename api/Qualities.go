package api

import (
	"github.com/Invoicing/error"
	"github.com/Invoicing/models"
	"github.com/labstack/echo"
	"strconv"
)

func ChangeQu2Des() echo.HandlerFunc {
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

		has, ok, err := models.ChangeQu2Des(id, count)
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

func ChangeQu2Car() echo.HandlerFunc {
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

		has, ok, err := models.ChangeQu2Car(id, count)
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

//通过订单获取某条采购信息
func GetQuByOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		orderId := c.FormValue("order_no")

		orderid, err := strconv.ParseInt(orderId, 10, 64)
		if err != nil {
			return sendError(errors.ID_ERROR, c)
		}

		pro, err := models.GetQuByOrder(orderid)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}

		return sendSuccess(1, pro, "以上为该单号的所有采购信息", c)
	}
}
