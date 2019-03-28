package api

import (
	"github.com/Invoicing/error"
	"github.com/Invoicing/models"
	"github.com/labstack/echo"
	"strconv"
)

//创建生产单
func CreateProOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		orderNo := c.FormValue("no")
		orderTag := c.FormValue("tag")
		//查询输入是否非法
		if orderNo == "" || orderTag == "" {
			return sendError(errors.INPUT_ERROR, c)
		}

		//查询该订单编号是否存在
		has, err := models.IsExistOrderNo(orderNo, PRODUCT)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		} //已存在
		if has {
			return sendError(errors.ORDER_EXIST, c)
		}

		err = models.CreateProOrder(orderNo, orderTag)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		return sendSuccess(1, "", "已经生成全部相关order", c)
	}
}

//向生产单插入零件
func InsertPro() echo.HandlerFunc {
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

		_, err = models.InsertProductComponet(orderNo, component_id, count)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}

		return sendSuccess(1, "", "添加零件成功", c)
	}
}

func ChangePro2Qu() echo.HandlerFunc {
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

		has, ok, err := models.ChangePro2Qu(id, count)
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
func GetProByOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		orderId := c.FormValue("order_no")

		orderid, err := strconv.ParseInt(orderId, 10, 64)
		if err != nil {
			return sendError(errors.ID_ERROR, c)
		}

		pro, err := models.GetProByOrder(orderid)
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}

		return sendSuccess(1, pro, "以上为该单号的所有采购信息", c)
	}
}
