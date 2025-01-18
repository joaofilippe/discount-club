package discountwebserver

import "github.com/labstack/echo/v4"

func (ws *WebServer) BuildRoutes() {
	ws.group.GET("/", func(c echo.Context) error {
		return c.JSON(200, struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    200,
			Message: "Discounts",
		})
	})
	ws.group.POST("/", ws.CreateDiscount)
}
