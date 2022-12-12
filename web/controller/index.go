package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (controller *Controller) index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", nil)
}
