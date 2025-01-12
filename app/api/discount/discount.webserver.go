package discountwebserver

import (
	"net/http"

	"github.com/joaofilippe/discount-club/app/domain/entities"
	"github.com/joaofilippe/discount-club/app/domain/iservices"
	"github.com/labstack/echo/v4"
)

type WebServer struct {
	server          *echo.Echo
	discountService iservices.IDiscount
}

func New(discountService iservices.IDiscount, server *echo.Echo) *WebServer {
	return &WebServer{
		server:          server,
		discountService: discountService,
	}
}

func (ws *WebServer) CreateDiscount(c echo.Context) error {
	request := new(CreateRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	validatedRequest, err := request.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	discount, err := entities.NewDiscount(
		validatedRequest.RestaurantID,
		validatedRequest.UserID,
		validatedRequest.Percentage,
		validatedRequest.StartDate,
		validatedRequest.EndDate,
		validatedRequest.Times,
		validatedRequest.Description,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	result, err := ws.discountService.Create(discount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}
