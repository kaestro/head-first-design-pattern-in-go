package factoryMethodPattern

import "fmt"

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type PizzaStore interface {
	CreatePizza(string) Pizza
	OrderPizza(string) Pizza
}

type PizzaBase struct {
	name string
}

func (p *PizzaBase) Prepare() {
	fmt.Println("Preparing " + p.name)
}

func (p *PizzaBase) Bake() {
	fmt.Println("Baking " + p.name)
}

func (p *PizzaBase) Cut() {
	fmt.Println("Cutting " + p.name)
}

func (p *PizzaBase) Box() {
	fmt.Println("Boxing " + p.name)
}

type NYStyleCheesePizza struct {
	PizzaBase
}

func (p *NYStyleCheesePizza) Prepare() {
	p.PizzaBase.Prepare()
}

func (p *NYStyleCheesePizza) Bake() {
	p.PizzaBase.Bake()
}

func (p *NYStyleCheesePizza) Cut() {
	p.PizzaBase.Cut()
}

func (p *NYStyleCheesePizza) Box() {
	p.PizzaBase.Box()
}

type NYStyleGreekPizza struct {
	PizzaBase
}

func (p *NYStyleGreekPizza) Prepare() {
	p.PizzaBase.Prepare()
}

func (p *NYStyleGreekPizza) Bake() {
	p.PizzaBase.Bake()
}

func (p *NYStyleGreekPizza) Cut() {
	p.PizzaBase.Cut()
}

func (p *NYStyleGreekPizza) Box() {
	p.PizzaBase.Box()
}

type NYStylePepperoniPizza struct {
	PizzaBase
}

func (p *NYStylePepperoniPizza) Prepare() {
	p.PizzaBase.Prepare()
}

func (p *NYStylePepperoniPizza) Bake() {
	p.PizzaBase.Bake()
}

func (p *NYStylePepperoniPizza) Cut() {
	p.PizzaBase.Cut()
}

func (p *NYStylePepperoniPizza) Box() {
	p.PizzaBase.Box()
}

type ChicagoStyleCheesePizza struct {
	PizzaBase
}

func (p *ChicagoStyleCheesePizza) Prepare() {
	p.PizzaBase.Prepare()
}

func (p *ChicagoStyleCheesePizza) Bake() {
	p.PizzaBase.Bake()
}

func (p *ChicagoStyleCheesePizza) Cut() {
	p.PizzaBase.Cut()
}

func (p *ChicagoStyleCheesePizza) Box() {
	p.PizzaBase.Box()
}

type ChicagoStyleGreekPizza struct {
	PizzaBase
}

func (p *ChicagoStyleGreekPizza) Prepare() {
	p.PizzaBase.Prepare()
}

func (p *ChicagoStyleGreekPizza) Bake() {
	p.PizzaBase.Bake()
}

func (p *ChicagoStyleGreekPizza) Cut() {
	p.PizzaBase.Cut()
}

func (p *ChicagoStyleGreekPizza) Box() {
	p.PizzaBase.Box()
}

type ChicagoStylePepperoniPizza struct {
	PizzaBase
}

func (p *ChicagoStylePepperoniPizza) Prepare() {
	p.PizzaBase.Prepare()
}

func (p *ChicagoStylePepperoniPizza) Bake() {
	p.PizzaBase.Bake()
}

func (p *ChicagoStylePepperoniPizza) Cut() {
	p.PizzaBase.Cut()
}

func (p *ChicagoStylePepperoniPizza) Box() {
	p.PizzaBase.Box()
}

type BasePizzaStore struct{}

func (bps *BasePizzaStore) OrderPizza(ps PizzaStore, pizzaType string) Pizza {
	pizza := ps.CreatePizza(pizzaType)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

type NYPizzaStore struct {
	BasePizzaStore
}

func (nps *NYPizzaStore) CreatePizza(pizzaType string) Pizza {
	var pizza Pizza
	if pizzaType == "cheese" {
		pizza = &NYStyleCheesePizza{PizzaBase: PizzaBase{name: "NY Style Cheese Pizza"}}
	} else if pizzaType == "greek" {
		pizza = &NYStyleGreekPizza{PizzaBase: PizzaBase{name: "NY Style Greek Pizza"}}
	} else if pizzaType == "pepperoni" {
		pizza = &NYStylePepperoniPizza{PizzaBase: PizzaBase{name: "NY Style Pepperoni Pizza"}}
	}
	return pizza
}

func (nps *NYPizzaStore) OrderPizza(pizzaType string) Pizza {
	return nps.BasePizzaStore.OrderPizza(nps, pizzaType)
}

type ChicagoPizzaStore struct {
	BasePizzaStore
}

func (cps *ChicagoPizzaStore) CreatePizza(pizzaType string) Pizza {
	var pizza Pizza
	if pizzaType == "cheese" {
		pizza = &ChicagoStyleCheesePizza{PizzaBase: PizzaBase{name: "Chicago Style Cheese Pizza"}}
	} else if pizzaType == "greek" {
		pizza = &ChicagoStyleGreekPizza{PizzaBase: PizzaBase{name: "Chicago Style Greek Pizza"}}
	} else if pizzaType == "pepperoni" {
		pizza = &ChicagoStylePepperoniPizza{PizzaBase: PizzaBase{name: "Chicago Style Pepperoni Pizza"}}
	}
	return pizza
}

func (cps *ChicagoPizzaStore) OrderPizza(pizzaType string) Pizza {
	return cps.BasePizzaStore.OrderPizza(cps, pizzaType)
}

func Simulate() {
	pizzaStores := []PizzaStore{&NYPizzaStore{}, &ChicagoPizzaStore{}}
	for _, ps := range pizzaStores {
		fmt.Println("Welcome to the Pizza Store!")
		ps.OrderPizza("cheese")
		ps.OrderPizza("greek")
		ps.OrderPizza("pepperoni")
		fmt.Println("")
	}
}
