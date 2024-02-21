package goodCoffeeShop

import "fmt"

const (
	EspressoDescription = "에스프레소"
	EspressoPrice       = 1.99

	HouseBlendDescription = "하우스 블렌드 커피"
	HouseBlendPrice       = 0.89

	MochaDescription = ", 모카"
	MochaPrice       = 0.5

	MilkDescription = ", 우유"
	MilkPrice       = 0.3

	SoyDescription = ", 두유"
	SoyPrice       = 0.2
)

type Beverage interface {
	Description() string
	Price() float64
}

type Espresso struct{}

func (e Espresso) Description() string {
	return EspressoDescription
}

func (e Espresso) Price() float64 {
	return EspressoPrice
}

type HouseBlend struct{}

func (hb HouseBlend) Description() string {
	return HouseBlendDescription
}

func (hb HouseBlend) Price() float64 {
	return HouseBlendPrice
}

type Mocha struct {
	Beverage
}

func (m Mocha) Description() string {
	return m.Beverage.Description() + MochaDescription
}

func (m Mocha) Price() float64 {
	return m.Beverage.Price() + MochaPrice
}

type Milk struct {
	Beverage
}

func (m Milk) Description() string {
	return m.Beverage.Description() + MilkDescription
}

func (m Milk) Price() float64 {
	return m.Beverage.Price() + MilkPrice
}

type Soy struct {
	Beverage
}

func (s Soy) Description() string {
	return s.Beverage.Description() + SoyDescription
}

func (s Soy) Price() float64 {
	return s.Beverage.Price() + SoyPrice
}

func SimulateGoodCoffeeShop() {
	beverages := make([]Beverage, 0, 5)

	beverages = append(beverages, Espresso{})
	beverages = append(beverages, HouseBlend{})
	beverages = append(beverages, Mocha{Espresso{}})
	beverages = append(beverages, Milk{HouseBlend{}})
	beverages = append(beverages, Soy{Milk{HouseBlend{}}})

	for _, beverage := range beverages {
		fmt.Println(beverage.Description(), beverage.Price())
	}
}
