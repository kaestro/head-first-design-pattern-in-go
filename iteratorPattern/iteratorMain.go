package main

import (
	"fmt"
	"iteratorPattern/menu"
	"iteratorPattern/waitress"
)

func main() {
	pancakeHouseMenu := menu.NewPancakeHouseMenu()
	dinerMenu := menu.NewDinerMenu()

	fmt.Println("Normal Waitress")
	normalWaitress := waitress.NewNormalWaitress(*pancakeHouseMenu, *dinerMenu)
	normalWaitress.PrintMenu()

	fmt.Println("\nIterator Waitress")
	iteratorWaitress := waitress.NewIteratorWaitress(*pancakeHouseMenu, *dinerMenu)
	iteratorWaitress.PrintMenu()
}
