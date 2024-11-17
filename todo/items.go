package todo

import "fmt"

type Item struct {
	Name    string
	Checked bool
}

func NewItem(index int) Item {
	return Item{
		Name: fmt.Sprintf("Item %d", index),
	}
}

func ItemList() []Item {
	items := []Item{}
	for i := 1; i < 10; i++ {
		items = append(items, NewItem(i))
	}
	return items
}
