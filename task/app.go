package task

import (
	"errors"
)

type App struct {
	Tasks []*Task
}

func NewApp() *App {
	return &App{
		Tasks: TaskList(),
	}
}

func (app *App) AddNewTask() *Task {
	t := NewTask(len(app.Tasks) + 1)
	app.Tasks = append(app.Tasks, t)
	return t
}

func (app *App) RemoveTask(id int) {
	i, err := app.findTask(id)
	if err == nil {
		app.Tasks = append(app.Tasks[:i], app.Tasks[i+1:]...)
	}
}

func (app *App) GetTask(id int) (*Task, error) {
	i, err := app.findTask(id)
	if err != nil {
		return nil, err
	}

	return app.Tasks[i], nil
}

func (app *App) findTask(id int) (int, error) {
	for index, t := range app.Tasks {
		if t.Id == id {
			return index, nil
		}
	}

	return -1, errors.New("Task not found")
}
