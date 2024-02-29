package waitress

import (
	"iteratorPattern/menu"
)

type IteratorWaitress struct {
	pancakeHouseMenu menu.PancakeHouseMenu
	dinerMenu        menu.DinerMenu
}

func NewIteratorWaitress(pancakeHouseMenu menu.PancakeHouseMenu, dinerMenu menu.DinerMenu) *IteratorWaitress {
	return &IteratorWaitress{pancakeHouseMenu, dinerMenu}
}

func (i *IteratorWaitress) PrintMenu() {
	breakFastItemsIterator := i.pancakeHouseMenu.CreateIterator()
	dinerItemsIterator := i.dinerMenu.CreateIterator()

	println("MENU\n----\nBREAKFAST")
	menu.PrintMenu(breakFastItemsIterator)

	println("\nLUNCH")
	menu.PrintMenu(dinerItemsIterator)
}
