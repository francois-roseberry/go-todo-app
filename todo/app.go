package todo

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
