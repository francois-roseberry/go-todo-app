package app

import (
	"context"
	"fmt"
	"strconv"

	"github.com/a-h/templ"
	"github.com/francois-roseberry/go-todo-app/internal/task"
	"github.com/francois-roseberry/go-todo-app/view/component"
	"github.com/francois-roseberry/go-todo-app/view/page"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

type TaskIdsPayload struct {
	TaskIds []int `form:"task-id"`
}

type Server struct {
	e    *echo.Echo
	port int
}

func NewServer(app *task.App, port int) *Server {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return render(c, 200, page.Index(app))
	})

	e.PUT("/status", func(c echo.Context) error {
		locked := c.QueryParam("locked")
		if locked == "true" {
			app.Locked = true
		} else {
			app.Locked = false
		}
		return render(c, 200, page.Container(app))
	})

	e.POST("/tasks", func(c echo.Context) error {
		task := app.AddNewTask()
		return render(c, 200, component.Task(task, app.Locked))
	})

	e.PUT("/tasks", func(c echo.Context) error {
		var payload TaskIdsPayload
		err := c.Bind(&payload)
		if err != nil {
			return err
		}
		return app.ReorderTasks(payload.TaskIds)
	})

	e.DELETE("/tasks", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.QueryParam("id"))
		app.RemoveTask(id)
		return nil
	})

	e.GET("/tasks/:id/edit-name", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		task, _ := app.GetTask(id)
		return render(c, 200, component.TaskNameEdit(task))
	})

	e.GET("/tasks/:id/display-name", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		task, _ := app.GetTask(id)
		return render(c, 200, component.TaskName(task, app.Locked))
	})

	e.PUT("/tasks/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		task, _ := app.GetTask(id)
		name := c.FormValue("task-name")
		if name != "" {
			task.Name = name
		}
		checked := c.FormValue("checked")
		if checked == "on" {
			task.Checked = true
		} else {
			task.Checked = false
		}
		return render(c, 200, component.TaskName(task, app.Locked))
	})

	return &Server{
		e:    e,
		port: port,
	}
}

func (s *Server) Start() {
	s.e.Logger.Fatal(s.e.Start(fmt.Sprintf(":%d", s.port)))
}

func (s *Server) Stop() {
	s.e.Shutdown(context.TODO())
}
