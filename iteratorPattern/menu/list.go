package menu

type List struct {
	items []interface{}
}

func (l *List) Add(item interface{}) {
	l.items = append(l.items, item)
}

func (l *List) Get(index int) interface{} {
	return l.items[index]
}

func (l *List) Remove(index int) {
	l.items = append(l.items[:index], l.items[index+1:]...)
}

func (l *List) Size() int {
	return len(l.items)
}
