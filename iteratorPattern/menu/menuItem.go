package menu

import "strconv"

type MenuItem struct {
	name        string
	description string
	vegetarian  bool
	price       float64
}

func NewMenuItem(name, description string, vegetarian bool, price float64) *MenuItem {
	return &MenuItem{name, description, vegetarian, price}
}

func (m *MenuItem) GetName() string {
	return m.name
}

func (m *MenuItem) GetDescription() string {
	return m.description
}

func (m *MenuItem) GetPrice() float64 {
	return m.price
}

func (m *MenuItem) IsVegetarian() bool {
	return m.vegetarian
}

func (m *MenuItem) String() string {
	return m.name + ", " + strconv.FormatFloat(m.price, 'f', 2, 64) + " -- " + m.description
}
