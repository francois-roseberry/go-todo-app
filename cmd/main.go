package main

import (
	"html/template"
	"io"
	"strconv"

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
	e.POST("/items", func(c echo.Context) error {
		item := app.AddNewItem()
		return c.Render(200, "item", item)
	})
	e.DELETE("/items", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.QueryParam("id"))
		app.RemoveItem(id)
		return nil
	})
	e.Logger.Fatal(e.Start(":3000"))
}
