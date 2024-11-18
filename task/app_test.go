package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppHasAnInitialListOfNineTasks(t *testing.T) {
	app := NewApp()
	assert.Equal(t, 9, len(app.Tasks), "App must initially have 9 tasks")
}

func TestAddTaskAtEndOfList(t *testing.T) {
	app := NewApp()
	l1 := len(app.Tasks)
	app.AddNewTask()
	l2 := len(app.Tasks)
	lastTask := app.Tasks[len(app.Tasks)-1]
	assert.Equal(t, l2, l1+1, "Appending new task must make list one task longer")
	assert.Equal(t, 10, lastTask.Id, "New task must have an id of 10")
	assert.False(t, lastTask.Checked, "New task must be unchecked")
}

func TestRemoveItem(t *testing.T) {
	idToRemove := 1
	app := NewApp()
	l1 := len(app.Tasks)
	app.RemoveTask(idToRemove)
	l2 := len(app.Tasks)
	assert.Equal(t, l2, l1-1, "Removing an item must make the list one item shorter")
}
