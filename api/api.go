package api

import (
	"github.com/Invoicing/error"
	"github.com/labstack/echo"
	"net/http"
	"unicode"
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
