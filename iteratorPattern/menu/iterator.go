package menu

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Menu interface {
	CreateIterator() Iterator
}

type PancakeHouseMenuIterator struct {
	items    *List
	position int
}

func NewPancakeHouseMenuIterator(items *List) *PancakeHouseMenuIterator {
	return &PancakeHouseMenuIterator{items: items}
}

func (p *PancakeHouseMenuIterator) HasNext() bool {
	if p.position >= len(p.items.items) || p.items.items[p.position] == nil {
		return false
	}
	return true
}

func (p *PancakeHouseMenuIterator) Next() interface{} {
	item := p.items.Get(p.position)
	p.position++
	return item
}

type DinerMenuIterator struct {
	items    []*MenuItem
	position int
}

func NewDinerMenuIterator(items []*MenuItem) *DinerMenuIterator {
	return &DinerMenuIterator{items: items}
}

func (d *DinerMenuIterator) HasNext() bool {
	if d.position >= len(d.items) || d.items[d.position] == nil {
		return false
	}
	return true
}

func (d *DinerMenuIterator) Next() interface{} {
	item := d.items[d.position]
	d.position++
	return item
}

func PrintMenu(menuIterator Iterator) {
	for menuIterator.HasNext() {
		item := menuIterator.Next().(*MenuItem)
		println(item.String())
	}
}
