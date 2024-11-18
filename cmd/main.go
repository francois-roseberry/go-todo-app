package main

import (
	"html/template"
	"io"
	"strconv"

	"github.com/francois-roseberry/go-todo-app/task"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	app := task.NewApp()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Renderer = newTemplate()
	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", *app)
	})
	e.POST("/tasks", func(c echo.Context) error {
		task := app.AddNewTask()
		return c.Render(200, "task", task)
	})
	e.DELETE("/tasks", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.QueryParam("id"))
		app.RemoveTask(id)
		return nil
	})
	e.Logger.Fatal(e.Start(":3000"))
}
