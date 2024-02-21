package badCoffeeShop

import "fmt"

type Beverage struct {
	Description string
}

func (b Beverage) GetDescription() string {
	return b.Description
}

type Espresso struct {
	Beverage
	Price float64
}

func (e Espresso) Cost() float64 {
	return e.Price
}

type HouseBlend struct {
	Beverage
	Price float64
}

func (hb HouseBlend) Cost() float64 {
	return hb.Price
}

type EspressoWithMocha struct {
	Beverage
	Price float64
}

func (e EspressoWithMocha) Cost() float64 {
	return e.Price
}

func SimulateBadCoffeeShop() {
	e := Espresso{Beverage{"에스프레소"}, 1.99}
	fmt.Println(e.GetDescription(), e.Cost())

	hb := HouseBlend{Beverage{"하우스 블렌드 커피"}, 0.89}
	fmt.Println(hb.GetDescription(), hb.Cost())

	em := EspressoWithMocha{Beverage{"에스프레소"}, 2.49}
	fmt.Println(em.GetDescription(), em.Cost())
}
