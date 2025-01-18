package api

import (
	"github.com/labstack/echo/v4"
	"github.com/joaofilippe/discount-club/app/application"
	"github.com/joaofilippe/discount-club/app/api/discount"
	
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
}