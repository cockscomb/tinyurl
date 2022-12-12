package controller

import (
	"github.com/cockscomb/tinyurl/usecase"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	tinyurl *usecase.TinyURLUsecase
}

func NewController(tinyurl *usecase.TinyURLUsecase) *Controller {
	return &Controller{tinyurl: tinyurl}
}

func (controller *Controller) Route(g *echo.Group) {
	g.GET("/", controller.index)
	g.POST("/generate", controller.generate)
	g.GET("/peek/:id", controller.peek)
	g.GET("/:id", controller.access)
}
