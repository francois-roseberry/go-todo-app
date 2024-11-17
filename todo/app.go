package todo

import (
	"errors"
)

type App struct {
	Items []Item
}

func NewApp() *App {
	return &App{
		Items: ItemList(),
	}
}

func (app *App) AddNewItem() Item {
	item := NewItem(len(app.Items) + 1)
	app.Items = append(app.Items, item)
	return item
}

func (app *App) RemoveItem(id int) {
	i, err := app.findItem(id)
	if err == nil {
		app.Items = append(app.Items[:i], app.Items[i+1:]...)
	}
}

func (app *App) findItem(id int) (int, error) {
	for index, item := range app.Items {
		if item.Id == id {
			return index, nil
		}
	}

	return -1, errors.New("Item not found")
}
