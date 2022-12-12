package web

import (
	"embed"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"html/template"
	"io"
	"net/url"
)

//go:embed template/*.html
var templates embed.FS

type TemplateConfig struct {
	Origin string `env:"ORIGIN,required"`
}

type Template struct {
	templates *template.Template
}

func NewTemplate(cfg *TemplateConfig) *Template {
	funcs := template.FuncMap{
		"absoluteURL": func(path string) (string, error) {
			u, err := url.Parse(cfg.Origin)
			if err != nil {
				return "", err
			}
			return u.JoinPath(path).String(), nil
		},
	}
	return &Template{
		templates: template.Must(template.New("").Funcs(funcs).ParseFS(templates, "template/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if data == nil {
		data = map[string]interface{}{}
	}
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["csrfToken"] = c.Get(middleware.DefaultCSRFConfig.ContextKey)
	}
	return t.templates.ExecuteTemplate(w, name, data)
}
