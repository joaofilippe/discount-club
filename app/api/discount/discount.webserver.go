package discountwebserver

import (
	"net/http"

	commonapi "github.com/joaofilippe/discount-club/app/api/common"
	"github.com/joaofilippe/discount-club/app/api/discount/requests"
	"github.com/joaofilippe/discount-club/app/domain/entities"
	"github.com/joaofilippe/discount-club/app/domain/iservices"
	"github.com/labstack/echo/v4"
)

type WebServer struct {
	group           *echo.Group
	discountService iservices.IDiscount
}

func New(discountService iservices.IDiscount, group *echo.Group) *WebServer {
	return &WebServer{
		group:           group.Group("/discounts"),
		discountService: discountService,
	}
}

func (ws *WebServer) CreateDiscount(c echo.Context) error {
	request := new(requests.CreateRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, commonapi.CommonInvalidResquest(err))
	}

	validatedRequest, err := request.Parse()
	if err != nil {
		return c.JSON(http.StatusBadRequest, commonapi.CommonInvalidResquest(err))
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

	discountDTO := NewDiscountDTOFromEntity(*result)

	return c.JSON(http.StatusCreated, commonapi.CommonResponse{
		Code:    http.StatusCreated,
		Message: "Discount created",
		Data:    discountDTO,
		Links: map[string]string{
			"self": "http://" + c.Request().Host + c.Request().RequestURI + "/",
			"get":  "http://" + c.Request().Host + c.Request().RequestURI + "/" + result.ID().String(),
		},
	})
}

func (ws *WebServer) GetDiscountByID(c echo.Context) error {
	id := c.Param("id")

	discount, err := ws.discountService.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, commonapi.CommonErrorResponse{
			Code:    http.StatusInternalServerError,
			ErrorMessage: err.Error(),
		})
	}

	discountDTO := NewDiscountDTOFromEntity(*discount)

	return c.JSON(http.StatusOK, commonapi.CommonResponse{
		Code:    http.StatusOK,
		Message: "Discount found",
		Data:    discountDTO,
	})
}

func (ws *WebServer) Verify(c echo.Context) error {
	code := c.Param("code")

	isValid, err := ws.discountService.Verify(code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, commonapi.CommonErrorResponse{
			Code:    http.StatusInternalServerError,
			ErrorMessage: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, commonapi.CommonResponse{
		Code:    http.StatusOK,
		Message: "Discount code is valid",
		Data: map[string]bool{
			"isValid": isValid,
		},
	})

}
