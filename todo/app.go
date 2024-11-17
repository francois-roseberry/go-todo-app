package todo

type App struct {
	Items []Item
}

func NewApp() *App {
	return &App{
		Items: ItemList(),
	}
}
