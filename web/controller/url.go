package controller

import (
	"errors"
	"github.com/cockscomb/tinyurl/domain/entity"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
)

func (controller *Controller) generate(c echo.Context) error {
	var param struct {
		URL string `form:"url"`
	}
	if err := c.Bind(&param); err != nil || param.URL == "" {
		return echo.ErrBadRequest
	}
	u, err := url.Parse(param.URL)
	if err != nil {
		return echo.ErrBadRequest
	}
	tinyURL, err := controller.tinyurl.Generate(c.Request().Context(), u)
	if err != nil {
		var validationErr *entity.ValidationError
		if errors.As(err, &validationErr) {
			return c.String(http.StatusBadRequest, validationErr.Error())
		}
		c.Logger().Error(err)
		return echo.ErrInternalServerError
	}
	return c.Redirect(http.StatusFound, "/peek/"+tinyURL.ID)
}

func (controller *Controller) peek(c echo.Context) error {
	var param struct {
		ID string `param:"id"`
	}
	if err := c.Bind(&param); err != nil {
		return echo.ErrBadRequest
	}
	tinyURL, err := controller.tinyurl.Peek(c.Request().Context(), param.ID)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return echo.ErrNotFound
		}
		c.Logger().Error(err)
		return echo.ErrInternalServerError
	}
	return c.Render(http.StatusOK, "peek.html", map[string]interface{}{
		"tinyurl": tinyURL,
	})
}

func (controller *Controller) access(c echo.Context) error {
	var param struct {
		ID string `param:"id"`
	}
	if err := c.Bind(&param); err != nil {
		return echo.ErrBadRequest
	}
	tinyURL, err := controller.tinyurl.Access(c.Request().Context(), param.ID)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			return echo.ErrNotFound
		}
		c.Logger().Error(err)
		return echo.ErrInternalServerError
	}
	return c.Redirect(http.StatusTemporaryRedirect, tinyURL.URL.String())
}
