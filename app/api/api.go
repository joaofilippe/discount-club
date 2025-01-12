package api

import (
	"github.com/joaofilippe/discount-club/app/application"
 )

type Api struct {
	application *application.Application
}

func New(application *application.Application) *Api {
	return &Api{
		application: application,
	}
}
