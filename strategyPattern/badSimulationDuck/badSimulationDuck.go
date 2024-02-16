package badSimulationDuck

import (
	"fmt"
)

type Duck interface {
	Quack()
	Swim()
	Display()
	fly()
}

type DuckBehavior struct{}

func (db DuckBehavior) Quack() {
	fmt.Println("quack")
}

func (db DuckBehavior) Swim() {
	fmt.Println("swim")
}

func (db DuckBehavior) fly() {
	fmt.Println("fly")
}

// DuckBehavior를 임베드하여 Duck 인터페이스를 구현하는 새로운 타입을 만들 수 있습니다.
// Display 메소드는 각 타입에서 구현해야 합니다.
type Mallard struct {
	DuckBehavior
}

func (m Mallard) Display() {
	// 구현해야 하는 메소드
	fmt.Println("I'm a Mallard Duck")
}

type Redhead struct {
	DuckBehavior
}

func (r Redhead) Display() {
	// 구현해야 하는 메소드
	fmt.Println("I'm a Redhead Duck")
}

type RubberDuck struct {
	DuckBehavior
}

func (r RubberDuck) Display() {
	// 구현해야 하는 메소드
	fmt.Println("I'm a Rubber Duck")
}

func Simulate() {
	var duck1 Duck = Mallard{}
	duck1.Quack()
	duck1.Swim()
	duck1.Display()
	duck1.fly()

	fmt.Println("")

	var duck2 Duck = Redhead{}
	duck2.Quack()
	duck2.Swim()
	duck2.Display()
	duck2.fly()

	fmt.Println("")

	var duck3 Duck = RubberDuck{}
	duck3.Quack()
	duck3.Swim()
	duck3.Display()
	duck3.fly()

	fmt.Println("")
}
