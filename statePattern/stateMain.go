package main

import (
	"fmt"
	"statePattern/badGumballMachine"
	"statePattern/goodGumballMachine"
)

func main() {
	fmt.Println("Bad Gumball Machine")
	badGumballMachine.Simulate()

	fmt.Println("\nGood Gumball Machine")
	goodGumballMachine.Simulate()
}
