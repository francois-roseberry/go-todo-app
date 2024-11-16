package todo

import "fmt"

type Item struct {
	Name    string
	Checked bool
}

func ItemList() []Item {
	items := []Item{}
	for i := 1; i < 10; i++ {
		items = append(items, Item{
			Name: fmt.Sprintf("Item %d", i),
		})
	}
	return items
}
