package menu

type DinerMenu struct {
	menuItems []*MenuItem
}

func NewDinerMenu() *DinerMenu {
	menuItems := make([]*MenuItem, 0)
	menuItems = append(menuItems, NewMenuItem("Vegetarian BLT", "(Fakin') Bacon with lettuce & tomato on whole wheat", true, 2.99))
	menuItems = append(menuItems, NewMenuItem("BLT", "Bacon with lettuce & tomato on whole wheat", false, 2.99))
	menuItems = append(menuItems, NewMenuItem("Soup of the day", "Soup of the day, with a side of potato salad", false, 3.29))
	menuItems = append(menuItems, NewMenuItem("Hotdog", "A hot dog, with saurkraut, relish, onions, topped with cheese", false, 3.05))

	return &DinerMenu{menuItems: menuItems}
}

func (d *DinerMenu) GetMenuItems() []*MenuItem {
	return d.menuItems
}

func (d *DinerMenu) AddItem(name, description string, vegetarian bool, price float64) {
	d.menuItems = append(d.menuItems, NewMenuItem(name, description, vegetarian, price))
}

func (d *DinerMenu) RemoveItem(index int) {
	d.menuItems = append(d.menuItems[:index], d.menuItems[index+1:]...)
}

func (d *DinerMenu) PrintMenu() {
	for _, item := range d.menuItems {
		println(item.String())
	}
}

func (d *DinerMenu) CreateIterator() Iterator {
	return NewDinerMenuIterator(d.menuItems)
}
