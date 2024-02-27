package dieBremerStadtmusikanten

type TurkeyAdapter struct {
	turkey Turkey
}

func (t *TurkeyAdapter) Quack() {
	t.turkey.Gobble()
}

func (t *TurkeyAdapter) Fly() {
	for i := 0; i < 5; i++ {
		t.turkey.Fly()
	}
}

// go style constructor. Go는 이렇게 일종의 팩토리 메소드를 만들어서 객체를 생성하는 경우가 많다.
func NewTurkeyAdapter(turkey Turkey) *TurkeyAdapter {
	return &TurkeyAdapter{turkey: turkey}
}
