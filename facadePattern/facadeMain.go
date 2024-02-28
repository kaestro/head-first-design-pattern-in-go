package main

import (
	"facadePattern/homeTheater"
	"fmt"
)

func main() {
	fmt.Println("Bad Home Theater")
	homeTheater.SimulateBadHomeTheater()

	fmt.Println("----------------------")
	fmt.Println("Good Home Theater")
	homeTheater.SimulateGoodHomeTheater()
}
