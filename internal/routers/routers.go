package router

import (
	"github.com/Paya-4970/Portfolio-v.3/internal/handlers"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"net/http"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func InitRouter(e *echo.Echo) {
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer
	e.Static("/static", "static")
	e.GET("/", handlers.HomeHandler)
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
}
