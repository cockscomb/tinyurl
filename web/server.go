package web

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
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

func NewServer(config *ServerConfig, db *dynamodb.Client) *Server {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		output, err := db.ListTables(c.Request().Context(), &dynamodb.ListTablesInput{})
		if err != nil {
			c.Logger().Error(err)
			return err
		}
		return c.JSON(http.StatusOK, output)
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
