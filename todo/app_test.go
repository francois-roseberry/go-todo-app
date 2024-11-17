package todo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppHasAnInitialListOfTenItems(t *testing.T) {
	app := NewApp()
	assert.Equal(t, 9, len(app.Items), "App must initially have 9 items")
}

func TestAddItemAtEndOfList(t *testing.T) {
	app := NewApp()
	l1 := len(app.Items)
	app.AddNewItem()
	l2 := len(app.Items)
	lastItem := app.Items[len(app.Items)-1]
	assert.Equal(t, l2, l1+1, "Appending new item must make list one item longer")
	assert.Equal(t, 10, lastItem.Id, "New item must have an id of 10")
	assert.False(t, lastItem.Checked, "New item must be unchecked")
}
