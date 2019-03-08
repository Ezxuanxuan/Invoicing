package api

import (
	"github.com/Invoicing/error"
	"github.com/Invoicing/models"
	"github.com/labstack/echo"
)

type OrderResult struct {
	OrderNo   string
	OrderType string
	Tag       string
}

func GetAllOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		orders, err := models.GetAllOrder()
		if err != nil {
			return sendError(errors.DO_ERROR, c)
		}
		if orders == nil {
			return sendError(errors.DO_ERROR, c)
		}
		results := make([]OrderResult, len(orders))
		for i, order := range orders {
			results[i].OrderNo = order.OrderNo
			switch order.OrderType {
			case IN:
				results[i].OrderType = "入库单"
				break
			case OUT:
				results[i].OrderType = "出库单"
				break
			case PURCHASE:
				results[i].OrderType = "采购单"
				break
			case PRODUCT:
				results[i].OrderType = "投产单"
				break
			case QUALITY:
				results[i].OrderType = "质检单"
				break
			case DESTROY:
				results[i].OrderType = "销毁单"
				break
			case CARRY:
				results[i].OrderType = "产成单"
				break
			}
			results[i].Tag = order.Tag
		}
		return sendSuccess(1, results, "以上为全部单号", c)
	}
}
