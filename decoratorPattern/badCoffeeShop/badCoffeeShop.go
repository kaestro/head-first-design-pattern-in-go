package badCoffeeShop

import "fmt"

const (
    EspressoDescription       = "에스프레소"
    EspressoPrice             = 1.99
    HouseBlendDescription     = "하우스 블렌드 커피"
    HouseBlendPrice           = 0.89
    EspressoWithMochaDescription = "에스프레소"
    EspressoWithMochaPrice       = 2.49
)

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

    espresso := Espresso{Description: EspressoDescription, Price: EspressoPrice}
    beverages = append(beverages, espresso)

    houseblend := HouseBlend{Description: HouseBlendDescription, Price: HouseBlendPrice}
    beverages = append(beverages, houseblend)

    espressoMocha := EspressoWithMocha{Description: EspressoWithMochaDescription, Price: EspressoWithMochaPrice}
    beverages = append(beverages, espressoMocha)

    for _, beverage := range beverages {
        fmt.Printf("%s: $%.2f\n", beverage.GetDescription(), beverage.Cost())
    }
}