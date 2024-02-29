package menu

type PancakeHouseMenu struct {
	menuItems *List
}

func NewPancakeHouseMenu() *PancakeHouseMenu {
	menuItems := &List{}
	menuItems.Add(NewMenuItem("K&B's Pancake Breakfast", "Pancakes with scrambled eggs, and toast", true, 2.99))
	menuItems.Add(NewMenuItem("Regular Pancake Breakfast", "Pancakes with fried eggs, sausage", false, 2.99))
	menuItems.Add(NewMenuItem("Blueberry Pancakes", "Pancakes made with fresh blueberries", true, 3.49))
	menuItems.Add(NewMenuItem("Waffles", "Waffles, with your choice of blueberries or strawberries", true, 3.59))

	return &PancakeHouseMenu{menuItems: menuItems}
}

func (p *PancakeHouseMenu) GetMenuItems() *List {
	return p.menuItems
}

func (p *PancakeHouseMenu) AddItem(name, description string, vegetarian bool, price float64) {
	p.menuItems.Add(NewMenuItem(name, description, vegetarian, price))
}

func (p *PancakeHouseMenu) RemoveItem(index int) {
	p.menuItems.Remove(index)
}

func (p *PancakeHouseMenu) PrintMenu() {
	iterator := p.CreateIterator()
	for iterator.HasNext() {
		item := iterator.Next().(*MenuItem)
		println(item.String())
	}
}

func (p *PancakeHouseMenu) CreateIterator() Iterator {
	return NewPancakeHouseMenuIterator(p.menuItems)
}
