package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"

	commonapi "github.com/joaofilippe/discount-club/app/api/common"
	"github.com/joaofilippe/discount-club/app/domain/entities"
)

func BuildCreateResponse(c echo.Context, user *entities.User) commonapi.CommonResponse {
	userDTO := NewUserDTOFromEntity(*user)
	baseURL := commonapi.GetBaseURL(c)

	return commonapi.CommonResponse{
		Code:    http.StatusCreated,
		Message: "User created",
		Data:    userDTO,
		Links: map[string]string{
			"self": baseURL + "/",
			"get":  baseURL + "/" + user.ID().String(),
		},
	}
}
