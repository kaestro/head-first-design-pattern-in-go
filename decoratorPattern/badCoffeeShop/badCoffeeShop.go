package badCoffeeShop

import "fmt"

type Beverage struct {
	Description string
	Price       float64
}

func (b Beverage) GetDescription() string {
	return b.Description
}

func (b Beverage) Cost() float64 {
	return b.Price
}

type Espresso = Beverage
type HouseBlend = Beverage
type EspressoWithMocha = Beverage

func SimulateBadCoffeeShop() {
	beverages := make([]Beverage, 0, 3)

	espresso := Espresso{Description: "에스프레소", Price: 1.99}
	beverages = append(beverages, espresso)

	houseblend := HouseBlend{Description: "하우스 블렌드 커피", Price: 0.89}
	beverages = append(beverages, houseblend)

	espressoMocha := EspressoWithMocha{Description: "에스프레소", Price: 2.49}
	beverages = append(beverages, espressoMocha)

	for _, beverage := range beverages {
		fmt.Printf("%s: $%.2f\n", beverage.GetDescription(), beverage.Cost())
	}
}
