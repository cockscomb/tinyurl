package web

import (
	"embed"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

//go:embed template/*.html
var templates embed.FS

type Template struct {
	templates *template.Template
}

func NewTemplate() *Template {
	return &Template{
		templates: template.Must(template.ParseFS(templates, "template/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
