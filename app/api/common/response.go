package commonapi

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type CommonResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    any               `json:"data"`
	Links   map[string]string `json:"_links"`
}

type CommonErrorResponse struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"error_message"`
	ErrorDetails any    `json:"error_details"`
}

func CommonInvalidResquest(err error) CommonErrorResponse {
	return CommonErrorResponse{
		Code:         http.StatusBadRequest,
		ErrorMessage: "Invalid request",
		ErrorDetails: err.Error(),
	}
}

func GetBaseURL(c echo.Context) string {
	return "http://" + c.Request().Host + c.Request().RequestURI
}
