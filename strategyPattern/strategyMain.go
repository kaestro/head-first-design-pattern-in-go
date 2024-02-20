package main

import (
	"fmt"
	"strategyPattern/badSimulationDuck"
	"strategyPattern/goodSimulationDuck"
)

func main() {
	fmt.Println("Bad Simulation Duck:")
	badSimulationDuck.Simulate()

	fmt.Println("Good Simulation Duck:")

	goodSimulationDuck.Simulate()
}
