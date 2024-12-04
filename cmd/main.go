package main

import (
	"github.com/francois-roseberry/go-todo-app/internal/app"
	"github.com/francois-roseberry/go-todo-app/internal/task"
)

func main() {
	s := app.NewServer(task.NewApp(), 3000)
	s.Start()
}
