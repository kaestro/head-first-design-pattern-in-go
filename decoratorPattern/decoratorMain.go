package main

import (
	"decoratorPattern/badCoffeeShop"
	"decoratorPattern/goodCoffeeShop"
	"fmt"
)

func main() {
	fmt.Println("Bad Coffee Shop")
	badCoffeeShop.SimulateBadCoffeeShop()

	fmt.Println("--------------------")

	fmt.Println("Good Coffee Shop")
	goodCoffeeShop.SimulateGoodCoffeeShop()
}
