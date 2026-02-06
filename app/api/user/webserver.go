package userwebserver

import (
	"net/http"

	commonapi "github.com/joaofilippe/discount-club/app/api/common"
	"github.com/joaofilippe/discount-club/app/api/user/requests"
	"github.com/joaofilippe/discount-club/app/api/user/responses"
	"github.com/joaofilippe/discount-club/app/domain/entities"
	usertype "github.com/joaofilippe/discount-club/app/domain/enums/user_type"
	"github.com/joaofilippe/discount-club/app/domain/iservices"
	"github.com/labstack/echo/v4"
)

type WebServer struct {
	group       *echo.Group
	userService iservices.IUser
}

func New(userService iservices.IUser, group *echo.Group) *WebServer {
	return &WebServer{
		group:       group.Group("/users"),
		userService: userService,
	}
}

func (ws *WebServer) CreateUser(c echo.Context) error {
	request := new(requests.CreateUserRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, commonapi.CommonInvalidResquest(err))
	}

	validatedRequest, err := request.Parse()
	if err != nil {
		return c.JSON(http.StatusBadRequest, commonapi.CommonInvalidResquest(err))
	}

	parsedUserType, ok := usertype.TryParse(validatedRequest.UserType)
	if !ok {
		return c.JSON(http.StatusBadRequest, commonapi.CommonErrorResponse{
			Code:         http.StatusBadRequest,
			ErrorMessage: "Invalid user type",
		})
	}

	user := entities.NewUser(
		validatedRequest.Email,
		validatedRequest.Password,
		parsedUserType,
	)

	result, err := ws.userService.Create(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, commonapi.CommonErrorResponse{
			Code:         http.StatusInternalServerError,
			ErrorMessage: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, responses.BuildCreateResponse(c, result))
}
