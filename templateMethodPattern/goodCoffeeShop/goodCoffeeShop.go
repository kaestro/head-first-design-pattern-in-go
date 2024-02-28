package goodCoffeeShop

import (
	"fmt"
)

type CaffeineBeverage interface {
	Brew()
	AddCondiments()
	BoilWater()
	PourInCup()
}

// Go는 이렇게 일부만 구현하고, 나머지는 더 하위 클래스에서 구현하도록 할 수 있다.
type CaffeineBeverageImpl struct{}

func (c *CaffeineBeverageImpl) BoilWater() {
	fmt.Println("Boiling water")
}

func (c *CaffeineBeverageImpl) PourInCup() {
	fmt.Println("Pouring into cup")
}

type Coffee struct {
	CaffeineBeverageImpl
}

func (c *Coffee) Brew() {
	fmt.Println("Dripping Coffee through filter")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("Adding Sugar and Milk")
}

type Tea struct {
	CaffeineBeverageImpl
}

func (t *Tea) Brew() {
	fmt.Println("Steeping the tea")
}

func (t *Tea) AddCondiments() {
	fmt.Println("Adding Lemon")
}

func PrepareRecipe(cb CaffeineBeverage) {
	cb.BoilWater()
	cb.Brew()
	cb.PourInCup()
	cb.AddCondiments()
}

func Simulate() {
	coffee := Coffee{}
	tea := Tea{}

	fmt.Println("Making Coffee...")
	PrepareRecipe(&coffee)

	fmt.Println("Making Tea...")
	PrepareRecipe(&tea)
}
