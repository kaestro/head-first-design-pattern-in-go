package main

import (
	"factoryPattern/badPizzaShop"
	"factoryPattern/factoryMethodPattern"
	"factoryPattern/simplePizzaFactory"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the Bad Pizza Shop!")
	badPizzaShop.Simulate()

	fmt.Println("")

	fmt.Println("Welcome to the Simple Pizza Factory Shop!")
	simplePizzaFactory.Simulate()

	fmt.Println("")

	fmt.Println("Welcome to the Factory Method Pattern Shop!")
	factoryMethodPattern.Simulate()
}
