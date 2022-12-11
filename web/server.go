package web

import (
	"context"
	"github.com/cockscomb/tinyurl/web/controller"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ServerConfig struct {
	Port int `env:"PORT" envDefault:"8080"`
}

type Server struct {
	e      *echo.Echo
	config *ServerConfig
}

func NewServer(config *ServerConfig, controller *controller.Controller) *Server {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	return &Server{
		e:      e,
		config: config,
	}
}

func (s *Server) Start() error {
	return s.e.Start(":" + strconv.Itoa(s.config.Port))
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}
