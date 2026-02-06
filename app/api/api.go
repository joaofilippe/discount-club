package api

import (
	discountwebserver "github.com/joaofilippe/discount-club/app/api/discount"
	userwebserver "github.com/joaofilippe/discount-club/app/api/user"
	"github.com/joaofilippe/discount-club/app/application"
	"github.com/labstack/echo/v4"
)

type Api struct {
	application application.IApplication
	server      *echo.Echo
}

func New(application application.IApplication, server *echo.Echo) *Api {
	return &Api{
		application: application,
		server:      server,
	}
}

func (a *Api) BuildRoutes() {
	apiV1 := a.server.Group("/api/v1")

	d := discountwebserver.New(a.application.DiscountService(), apiV1)
	d.BuildRoutes()

	u := userwebserver.New(a.application.UserService(), apiV1)
	u.ConfigureRoutes()
}
