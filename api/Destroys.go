package api

import (
	"github.com/Invoicing/error"
	"github.com/Invoicing/models"
	"github.com/labstack/echo"
	"strconv"
)

//通过订单获取某条采购信息
func GetDesByOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		orderId := c.FormValue("order_no")

		orderid, err := strconv.ParseInt(orderId, 10, 64)
		if err != nil {
			return sendError(errors.ID_ERROR, c)
		}

		pro, err := models.GetDesByOrder(orderid)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}

		return sendSuccess(1, pro, "以上为该单号的所有采购信息", c)
	}
}
