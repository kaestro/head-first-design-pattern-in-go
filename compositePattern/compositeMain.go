package main

import (
	"compositePattern/menu"
	"fmt"
)

func main() {
	fmt.Println("Composite Pattern")

	pancakeHouseMenu := menu.NewMenu("PANCAKE HOUSE MENU", "Breakfast")
	dinerMenu := menu.NewMenu("DINER MENU", "Lunch")
	cafeMenu := menu.NewMenu("CAFE MENU", "Dinner")

	waitress := menu.NewMenu("ALL MENUS", "All menus combined")

	waitress.Add(pancakeHouseMenu)
	waitress.Add(dinerMenu)
	waitress.Add(cafeMenu)

	pancakeHouseMenu.Add(menu.NewMenuItem("K&B's Pancake Breakfast", "Pancakes with scrambled eggs, and toast", 2.99))
	pancakeHouseMenu.Add(menu.NewMenuItem("Regular Pancake Breakfast", "Pancakes with fried eggs, sausage", 2.99))
	pancakeHouseMenu.Add(menu.NewMenuItem("Blueberry Pancakes", "Pancakes made with fresh blueberries", 3.49))

	dinerMenu.Add(menu.NewMenuItem("Vegetarian BLT", "(Fakin') Bacon with lettuce & tomato on whole wheat", 2.99))
	dinerMenu.Add(menu.NewMenuItem("BLT", "Bacon with lettuce & tomato on whole wheat", 2.99))

	cafeMenu.Add(menu.NewMenuItem("Veggie Burger and Air Fries", "Veggie burger on a whole wheat bun, lettuce, tomato, and fries", 3.29))

	waitress.Print()
}
