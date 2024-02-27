package dieBremerStadtmusikanten

type Turkey interface {
	Gobble()
	Fly()
}

type WildTurkey struct {
}

func (w *WildTurkey) Gobble() {
	println("Gobble gobble")
}

func (w *WildTurkey) Fly() {
	println("I'm flying a short distance")
}
