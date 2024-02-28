package badCoffeeShop

import (
	"fmt"
)

type CaffeineBeverage interface {
	PrepareRecipe()
	BoilWater()
	PourInCup()
}

type Coffee struct {
	CaffeineBeverage
}

func (c *Coffee) PrepareRecipe() {
	fmt.Println("Preparing Coffee")

	c.BoilWater()
	c.BrewCoffeeGrinds()
	c.PourInCup()
	c.AddSugarAndMilk()
}

func (c *Coffee) BoilWater() {
	fmt.Println("Boiling water")
}

func (c *Coffee) BrewCoffeeGrinds() {
	fmt.Println("Dripping Coffee through filter")
}

func (c *Coffee) PourInCup() {
	fmt.Println("Pouring into cup")
}

func (c *Coffee) AddSugarAndMilk() {
	fmt.Println("Adding Sugar and Milk")
}

type Tea struct {
	CaffeineBeverage
}

func (t *Tea) PrepareRecipe() {
	fmt.Println("Preparing Tea")

	t.BoilWater()
	t.SteepTeaBag()
	t.PourInCup()
	t.AddLemon()
}

func (t *Tea) BoilWater() {
	fmt.Println("Boiling water")
}

func (t *Tea) SteepTeaBag() {
	fmt.Println("Steeping the tea")
}

func (t *Tea) PourInCup() {
	fmt.Println("Pouring into cup")
}

func (t *Tea) AddLemon() {
	fmt.Println("Adding Lemon")
}

func Simulate() {
	coffee := Coffee{}
	coffee.PrepareRecipe()

	tea := Tea{}
	tea.PrepareRecipe()
}
