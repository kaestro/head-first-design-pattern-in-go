package menu

import (
	"fmt"
	"strconv"
)

type MenuComponent interface {
	Add(MenuComponent)
	Remove(MenuComponent)
	GetChild(int) MenuComponent
	GetName() string
	GetDescription() string
	GetPrice() float64
	Print()
}

type MenuNotSupported struct {
}

func (m *MenuNotSupported) Add(MenuComponent) {
	panic("Unsupported Operation")
}

func (m *MenuNotSupported) Remove(MenuComponent) {
	panic("Unsupported Operation")
}

func (m *MenuNotSupported) GetChild(int) MenuComponent {
	panic("Unsupported Operation")
}

type MenuItem struct {
	MenuNotSupported
	name        string
	description string
	price       float64
}

func NewMenuItem(name, description string, price float64) *MenuItem {
	return &MenuItem{name: name, description: description, price: price}
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

func (m *MenuItem) Print() {
	fmt.Println(m.GetName() + ", " + m.GetDescription() + ", " + strconv.FormatFloat(m.GetPrice(), 'f', 2, 64))
}

type Menu struct {
	MenuComponent
	menuComponents []MenuComponent
	name           string
	description    string
}

func NewMenu(name, description string) *Menu {
	return &Menu{name: name, description: description}
}

func (m *Menu) Add(menuComponent MenuComponent) {
	m.menuComponents = append(m.menuComponents, menuComponent)
}

func (m *Menu) Remove(menuComponent MenuComponent) {
	for i, c := range m.menuComponents {
		if c == menuComponent {
			m.menuComponents = append(m.menuComponents[:i], m.menuComponents[i+1:]...)
		}
	}
}

func (m *Menu) GetChild(i int) MenuComponent {
	return m.menuComponents[i]
}

func (m *Menu) GetName() string {
	return m.name
}

func (m *Menu) GetDescription() string {
	return m.description
}

func (m *Menu) Print() {
	fmt.Println(m.GetName() + ", " + m.GetDescription())
	fmt.Println("---------------------")

	for _, c := range m.menuComponents {
		c.Print()
	}
}
