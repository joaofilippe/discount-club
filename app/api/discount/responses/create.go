package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"

	commonapi "github.com/joaofilippe/discount-club/app/api/common"
	"github.com/joaofilippe/discount-club/app/domain/entities"
)

func BuildCreateResponse(c echo.Context, discount *entities.Discount) commonapi.CommonResponse {
	discountDTO := NewDiscountDTOFromEntity(*discount)

	return commonapi.CommonResponse{
		Code:    http.StatusCreated,
		Message: "Discount created",
		Data:    discountDTO,
		Links: map[string]string{
			"self": "http://" + c.Request().Host + c.Request().RequestURI + "/",
			"get":  "http://" + c.Request().Host + c.Request().RequestURI + "/" + discount.ID().String(),
		},
	}
}
