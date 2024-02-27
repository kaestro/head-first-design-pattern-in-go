package dieBremerStadtmusikanten

type Duck interface {
	Quack()
	Fly()
}

type MallardDuck struct {
}

func (m *MallardDuck) Quack() {
	println("Quack")
}

func (m *MallardDuck) Fly() {
	println("I'm flying")
}

type RocketDuck struct {
}

func (r *RocketDuck) Quack() {
	println("Boom!")
}

func (r *RocketDuck) Fly() {
	println("i'm flying with a rocket!")
}
