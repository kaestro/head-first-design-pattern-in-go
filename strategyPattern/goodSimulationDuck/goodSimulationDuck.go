package goodSimulationDuck

import "fmt"

type FlyBehavior interface {
	fly()
}

type QuackBehavior interface {
	quack()
}

type FlyWithWings struct{}

func (fww FlyWithWings) fly() {
	fmt.Println("flying With Wings")
}

type FlyNoWay struct{}

func (fnw FlyNoWay) fly() {
	fmt.Println("I can't fly")
}

type Quack struct{}

func (q Quack) quack() {
	fmt.Println("quack")
}

type Squeak struct{}

type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (s Squeak) quack() {
	fmt.Println("squeak")
}

type MuteQuack struct{}

func (mq MuteQuack) quack() {
	fmt.Println("<< Silence >>")
}

func Simulate() {
	mallardDuck := Duck{FlyWithWings{}, Quack{}}
	mallardDuck.flyBehavior.fly()
	mallardDuck.quackBehavior.quack()

	modelDuck := Duck{FlyNoWay{}, Quack{}}
	modelDuck.flyBehavior.fly()
	modelDuck.quackBehavior.quack()
}
