package api

import (
	"github.com/Invoicing/error"
	"github.com/Invoicing/models"
	"github.com/labstack/echo"
	"strconv"
)

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

func GetPurchasesByOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		orderId := c.FormValue("order_id")

		orderid, err := strconv.ParseInt(orderId, 10, 64)
		if err != nil {
			return sendError(errors.ID_ERROR, c)
		}

		purchases, has, err := models.GetPurchasesByOrder(orderid)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if !has {
			return sendError(errors.Order_NOT_EXIST, c)
		}

		return sendSuccess(1, purchases, "以上为该单号的所有采购信息", c)
	}
}

func CreatePurchaseOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		ordeNo := c.FormValue("no")
		orderTag := c.FormValue("tag")
		//查询输入是否非法
		if ordeNo == "" || orderTag == "" {
			return sendError(errors.INPUT_ERROR, c)
		}

		//查询该订单编号是否存在
		has, err := models.IsExistOrderNo(ordeNo, PURCHASE)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		} //已存在
		if has {
			return sendError(errors.ORDER_EXIST, c)
		}

		err = models.CreateAllOrder(ordeNo, orderTag)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, "", "已经生成全部相关order", c)
	}
}

//func ChangePurchaseById() echo.HandlerFunc {
//	return func(c echo.Context) error {
//		id := c.FormValue("id")
//
//		Id, err := strconv.ParseInt(id, 10, 64)
//		if err != nil {
//			return sendError(errors.INPUT_ERROR, c)
//		}
//
//	}
//}
