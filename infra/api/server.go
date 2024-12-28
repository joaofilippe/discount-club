package api

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	server *echo.Echo
}

var (
	instance *Server
	once     sync.Once
)

func NewServer() *Server {
	once.Do(func() {
		e := echo.New()
		e.Use(middleware.Logger())
		e.Use(middleware.Recover())

		instance = &Server{
			server: e,
		}

		instance.buildRoutes()
	})

	return instance
}

func (s *Server) Start(address string) error {
	return s.server.Start(address)
}

func (s *Server) buildRoutes() {
	s.server.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	// Add other routes here
}
