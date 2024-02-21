package goodCoffeeShop

import "fmt"

type Beverage interface {
	GetDescription() string
	Cost() float64
}

type Espresso struct {
	Description string
}

func (e Espresso) GetDescription() string {
	return e.Description
}

func (e Espresso) Cost() float64 {
	return 1.99
}

type HouseBlend struct {
	Description string
}

func (hb HouseBlend) GetDescription() string {
	return hb.Description
}

func (hb HouseBlend) Cost() float64 {
	return 0.89
}

type Mocha struct {
	Beverage
}

func (m Mocha) GetDescription() string {
	return m.Beverage.GetDescription() + ", 모카"
}

func (m Mocha) Cost() float64 {
	return m.Beverage.Cost() + 0.5
}

type Milk struct {
	Beverage
}

func (m Milk) GetDescription() string {
	return m.Beverage.GetDescription() + ", 우유"
}

func (m Milk) Cost() float64 {
	return m.Beverage.Cost() + 0.3
}

type Soy struct {
	Beverage
}

func (s Soy) GetDescription() string {
	return s.Beverage.GetDescription() + ", 두유"
}

func (s Soy) Cost() float64 {
	return s.Beverage.Cost() + 0.2
}

func SimulateGoodCoffeeShop() {
	e := Espresso{"에스프레소"}
	fmt.Println(e.GetDescription(), e.Cost())

	hb := HouseBlend{"하우스 블렌드 커피"}
	fmt.Println(hb.GetDescription(), hb.Cost())

	em := Mocha{e}
	fmt.Println(em.GetDescription(), em.Cost())

	hbm := Milk{hb}
	fmt.Println(hbm.GetDescription(), hbm.Cost())

	hbms := Soy{hbm}
	fmt.Println(hbms.GetDescription(), hbms.Cost())
}
