package waitress

import (
	"iteratorPattern/menu"
)

type NormalWaitress struct {
	pancakeHouseMenu menu.PancakeHouseMenu
	dinerMenu        menu.DinerMenu
}

func NewNormalWaitress(pancakeHouseMenu menu.PancakeHouseMenu, dinerMenu menu.DinerMenu) *NormalWaitress {
	return &NormalWaitress{pancakeHouseMenu, dinerMenu}
}

func (n *NormalWaitress) PrintMenu() {
	breakFastItems := n.pancakeHouseMenu.GetMenuItems()
	dinerItems := n.dinerMenu.GetMenuItems()

	println("MENU\n----\nBREAKFAST")
	for i := 0; i < breakFastItems.Size(); i++ {
		item := breakFastItems.Get(i).(*menu.MenuItem)
		println(item.String())
	}

	println("\nLUNCH")
	for _, item := range dinerItems {
		println(item.String())
	}
}
