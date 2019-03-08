package api

import (
	"github.com/Invoicing/error"
	"github.com/labstack/echo"
	"net/http"
	"unicode"
)

//订单类型
const (
	IN       = 1
	OUT      = 2
	PURCHASE = 3
	PRODUCT  = 4
	QUALITY  = 5
	DESTROY  = 6
	CARRY    = 7
)

func sendError(restful *errors.Restful, c echo.Context) error {
	return c.JSONPretty(http.StatusOK, restful, "      ")
}

func sendSuccess(code int, data interface{}, msg string, c echo.Context) error {
	successful := &errors.Successful{code, data, msg}
	return c.JSONPretty(http.StatusOK, successful, "      ")
}

//判断字符串是否为纯数字
func isAllNumic(s string) bool {
	for i := 0; i < len(s); i++ {
		if !unicode.IsNumber(rune(s[i])) {
			return false
		}
	}
	return true
}
