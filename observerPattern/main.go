package main

import (
	"fmt"
	goodObserver "observerPattern/GoodObserver"
	observerByPull "observerPattern/ObserverByPull"
	badObserver "observerPattern/badObserver"
)

func main() {
	fmt.Println("Bad Observer Pattern")
	badObserver.SimulateWeatherData()

	fmt.Println("\nGood Observer Pattern")
	goodObserver.SimulateWeatherData()

	fmt.Println("\nObserver By Pull Pattern")
	observerByPull.SimulateWeatherData()
}
