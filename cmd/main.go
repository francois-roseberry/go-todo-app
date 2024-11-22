package main

import (
	"strconv"

	"github.com/a-h/templ"
	"github.com/francois-roseberry/go-todo-app/internal/task"
	"github.com/francois-roseberry/go-todo-app/view/component"
	"github.com/francois-roseberry/go-todo-app/view/page"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

func main() {
	app := task.NewApp()

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return Render(c, 200, page.Index(app))
	})
	e.POST("/tasks", func(c echo.Context) error {
		task := app.AddNewTask()
		return Render(c, 200, component.Task(task))
	})
	e.DELETE("/tasks", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.QueryParam("id"))
		app.RemoveTask(id)
		return nil
	})
	e.GET("/tasks/:id/edit-name", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		task, _ := app.GetTask(id)
		return Render(c, 200, component.TaskNameEdit(task))
	})
	e.GET("/tasks/:id/display-name", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		task, _ := app.GetTask(id)
		return Render(c, 200, component.TaskName(task))
	})
	e.PUT("/tasks/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		name := c.FormValue("task-name")
		// TODO validate name
		task, _ := app.GetTask(id)
		task.Name = name
		return Render(c, 200, component.TaskName(task))
	})
	e.Logger.Fatal(e.Start(":3000"))
}
