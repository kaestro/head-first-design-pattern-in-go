package main

import (
	"fmt"
	"templateMethodPattern/badCoffeeShop"
	"templateMethodPattern/goodCoffeeShop"
)

func main() {
	fmt.Println("badCoffeeShop")
	fmt.Println("--------------")
	badCoffeeShop.Simulate()
	fmt.Println("--------------")

	fmt.Println("goodCoffeeShop")
	fmt.Println("--------------")
	goodCoffeeShop.Simulate()
	fmt.Println("--------------")
}
