package discountwebserver

import (
	"net/http"

	"github.com/joaofilippe/discount-club/app/domain/entities"
	"github.com/joaofilippe/discount-club/app/domain/iservices"
	"github.com/labstack/echo/v4"
)

type WebServer struct {
	discountService iservices.IDiscount
}

func New(discountService iservices.IDiscount) *WebServer {
	return &WebServer{
		discountService: discountService,
	}
}

func (ws *WebServer) CreateDiscount(c echo.Context) error {
	discount := new(entities.Discount)
	if err := c.Bind(discount); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	discount, err := ws.discountService.Create(discount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, discount)
}