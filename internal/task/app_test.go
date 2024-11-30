package task

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	t.Run("given new app", func(t *testing.T) {
		app := NewApp()
		t.Run("has an initial list of 9 tasks", func(t *testing.T) {
			assert.Equal(t, 9, len(app.Tasks))
		})
	})
}

func TestAddTask(t *testing.T) {
	t.Run("given task to add at end of list", func(t *testing.T) {
		app := NewApp()
		l1 := len(app.Tasks)
		app.AddNewTask()
		l2 := len(app.Tasks)
		lastTask := app.Tasks[len(app.Tasks)-1]
		t.Run("then task list is one item longer", func(t *testing.T) {
			assert.Equal(t, l2, l1+1)
		})
		t.Run("then new task has an id of 10", func(t *testing.T) {
			assert.Equal(t, 10, lastTask.Id)
		})
		t.Run("then new task must be unchecked", func(t *testing.T) {
			assert.False(t, lastTask.Checked)
		})
	})
}

func TestRemoveTask(t *testing.T) {
	t.Run("given task id to remove does not exist", func(t *testing.T) {
		idToRemove := 111
		app := NewApp()
		l1 := len(app.Tasks)
		app.RemoveTask(idToRemove)
		l2 := len(app.Tasks)
		t.Run("then does nothing", func(t *testing.T) {
			assert.Equal(t, l2, l1)
		})
	})
	t.Run("given task id to remove that exists", func(t *testing.T) {
		idToRemove := 1
		app := NewApp()
		l1 := len(app.Tasks)
		app.RemoveTask(idToRemove)
		l2 := len(app.Tasks)
		t.Run("then task list is one item shorter", func(t *testing.T) {
			assert.Equal(t, l2, l1-1)
		})
	})
}

func TestGetTask(t *testing.T) {
	t.Run("given id that does not exist", func(t *testing.T) {
		app := NewApp()
		_, err := app.GetTask(11)
		t.Run("then returns error", func(t *testing.T) {
			assert.Equal(t, errors.New("task not found"), err)
		})
	})

	t.Run("given id that exists", func(t *testing.T) {
		app := NewApp()
		task, err := app.GetTask(1)
		t.Run("then returns no error", func(t *testing.T) {
			assert.Nil(t, err)
		})

		t.Run("then returns requested task", func(t *testing.T) {
			assert.Equal(t, 1, task.Id)
		})
	})
}

func TestReorderTasks(t *testing.T) {
	t.Run("given list of ids of different size than task list", func(t *testing.T) {
		app := NewApp()
		newOrder := []int{1, 2, 3}
		app.ReorderTasks(newOrder)
		t.Run("then returns error", func(t *testing.T) {
			assert.Equal(t, errors.New("provided taskIds must be the same length as app tasks"), app.ReorderTasks(newOrder))
		})
	})
	t.Run("given list of ids in same order", func(t *testing.T) {
		app := NewApp()
		newOrder := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		_ = app.ReorderTasks(newOrder)
		t.Run("then nothing changes", func(t *testing.T) {
			assert.Equal(t, newOrder, transform(app.Tasks, func(t *Task) int {
				return t.Id
			}))
		})
	})

	t.Run("given list of mixed up ids", func(t *testing.T) {
		app := NewApp()
		newOrder := []int{4, 2, 1, 3, 5, 6, 7, 8, 9}
		_ = app.ReorderTasks(newOrder)
		t.Run("then tasks are re-ordered", func(t *testing.T) {
			assert.Equal(t, newOrder, transform(app.Tasks, func(t *Task) int {
				return t.Id
			}))
		})
	})
}

func transform[U any, V any](arr []U, fn func(U) V) (res []V) {
	for _, elem := range arr {
		res = append(res, fn(elem))
	}
	return
}
