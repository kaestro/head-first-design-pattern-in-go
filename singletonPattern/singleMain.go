package main

import (
	"fmt"
	"singletonPattern/badChocolateBoiler"
	"singletonPattern/goodChocolateBoiler"
)

func main() {
	fmt.Println("badChocolateBoiler")
	fmt.Println("===============\n")
	badChocolateBoiler.Simulate()

	fmt.Println("===============\n")
	fmt.Println("goodChocolateBoiler")
	goodChocolateBoiler.Simulate()
}
