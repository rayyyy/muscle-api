package handler

import (
	"github.com/labstack/echo/v4"
)

// APIError エラー時のレスポンス
type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// JSONErrorHandler error responce
func JSONErrorHandler(err error, c echo.Context) {
	var code int
	var message string

	if httpError, ok := err.(*echo.HTTPError); ok {
		code = httpError.Code
		switch code {
		case 401:
			message = "未ログインです。"
		case 403:
			message = "権限がありません。"
		case 404:
			message = "存在しません。"

		default:
			code = 500
			message = "サーバーエラーが起きました。"
		}
	}

	c.JSON(code, APIError{
		Status:  code,
		Message: message,
	})
}
