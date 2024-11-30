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

// Re-order the task list to follow an array of ids passed in
// Assumes that:
// - array contains unique values
// - array contains the ids of the existing tasks
func (app *App) ReorderTasks(taskIds []int) error {
	if len(taskIds) != len(app.Tasks) {
		return errors.New("provided taskIds must be the same length as app tasks")
	}
	tasks := make([]*Task, len(app.Tasks))
	for i, id := range taskIds {
		t, _ := app.GetTask(id)
		tasks[i] = t
	}
	app.Tasks = tasks
	return nil
}

func (app *App) findTask(id int) (int, error) {
	for index, t := range app.Tasks {
		if t.Id == id {
			return index, nil
		}
	}

	return -1, errors.New("task not found")
}
