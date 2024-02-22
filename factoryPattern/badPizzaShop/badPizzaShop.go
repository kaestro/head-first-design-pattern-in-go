package badPizzaShop

import "fmt"

type PizzaStore struct {
	pizza Pizza
}

func (ps *PizzaStore) OrderPizza(pizzaType string) Pizza {
	if pizzaType == "cheese" {
		ps.pizza = &CheesePizza{BasePizza: BasePizza{name: "Cheese Pizza"}}
	} else if pizzaType == "pepperoni" {
		ps.pizza = &PepperoniPizza{BasePizza: BasePizza{name: "Pepperoni Pizza"}}
	} else if pizzaType == "clam" {
		ps.pizza = &ClamPizza{BasePizza: BasePizza{name: "Clam Pizza"}}
	} else if pizzaType == "veggie" {
		ps.pizza = &VeggiePizza{BasePizza: BasePizza{name: "Veggie Pizza"}}
	}

	ps.pizza.Prepare()
	ps.pizza.Bake()
	ps.pizza.Cut()
	ps.pizza.Box()
	return ps.pizza
}

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
	GetName() string
}

type BasePizza struct {
	name string
}

func (p *BasePizza) Prepare() {
	fmt.Println("Preparing " + p.name)
}

func (p *BasePizza) Bake() {
	fmt.Println("Baking " + p.name)
}

func (p *BasePizza) Cut() {
	fmt.Println("Cutting " + p.name)
}

func (p *BasePizza) Box() {
	fmt.Println("Boxing " + p.name)
}

func (p *BasePizza) GetName() string {
	return p.name
}

type CheesePizza struct {
	BasePizza
}

type PepperoniPizza struct {
	BasePizza
}

type ClamPizza struct {
	BasePizza
}

type VeggiePizza struct {
	BasePizza
}

func Simulate() {
	pizzaStore := PizzaStore{}
	pizzaStore.OrderPizza("cheese")
}
