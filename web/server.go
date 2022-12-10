package web

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
	e *echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	return &Server{
		e: e,
	}
}

func (s *Server) Start(address string) error {
	return s.e.Start(address)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}
