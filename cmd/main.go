package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/francois-roseberry/go-todo-app/todo"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.Must(template.ParseGlob("views/components/*.html")).ParseGlob("views/pages/*.html")),
	}
}

func main() {
	app := todo.NewApp()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", *app)
	})
	e.Logger.Fatal(e.Start(":3000"))
}
