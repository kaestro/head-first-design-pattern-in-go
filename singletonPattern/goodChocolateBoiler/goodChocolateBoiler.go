package goodChocolateBoiler

import (
	"fmt"
	"sync"
)

type ChocolateBoiler struct {
	empty  bool
	boiled bool
}

var instance *ChocolateBoiler
var once sync.Once

func GetInstance() *ChocolateBoiler {
	once.Do(func() {
		instance = &ChocolateBoiler{empty: true, boiled: false}
	})
	return instance
}

// 우유/초콜릿을 채워넣는다.
func (cb *ChocolateBoiler) Fill() {
	if cb.empty {
		cb.empty = false
		cb.boiled = false
		fmt.Println("우유/초콜릿을 채워넣는다.")
	} else {
		fmt.Println("이미 채워진 보일러에는 더 이상 채울 수 없습니다.")
	}
}

// 끓인 우유/초콜릿을 다음 단계로 넘긴다.
func (cb *ChocolateBoiler) Drain() {
	if !cb.empty && cb.boiled {
		cb.empty = true
		fmt.Println("끓인 우유/초콜릿을 다음 단계로 넘긴다.")
	} else {
		fmt.Println("끓지 않은 우유/초콜릿은 다음 단계로 넘길 수 없습니다.")
	}
}

// 우유/초콜릿을 끓인다.
func (cb *ChocolateBoiler) Boil() {
	if !cb.empty && !cb.boiled {
		cb.boiled = true
		fmt.Println("우유/초콜릿을 끓인다.")
	} else {
		fmt.Println("비어있는 보일러는 끓일 수 없습니다.")
	}
}

func (cb *ChocolateBoiler) IsEmpty() bool {
	return cb.empty
}

func (cb *ChocolateBoiler) IsBoiled() bool {
	return cb.boiled
}

func Simulate() {
	cb1 := GetInstance()
	cb2 := GetInstance()

	cb1.Fill()
	cb2.Boil()
	cb1.Drain()
}
