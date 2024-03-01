package badGumballMachine

import (
	"fmt"
	"strconv"
)

type GumballMachine struct {
	state int
	count int
}

const (
	SoldOut    = 0
	NoQuarter  = 1
	HasQuarter = 2
	Sold       = 3
)

func NewGumballMachine(count int) *GumballMachine {
	if count > 0 {
		return &GumballMachine{state: NoQuarter, count: count}
	} else {
		return &GumballMachine{state: SoldOut, count: count}
	}
}

func (m *GumballMachine) InsertQuarter() {
	switch m.state {
	case HasQuarter:
		println("You can't insert another quarter")
	case NoQuarter:
		m.state = HasQuarter
		println("You inserted a quarter")
	case SoldOut:
		println("You can't insert a quarter, the machine is sold out")
	case Sold:
		println("Please wait, we're already giving you a gumball")
	}
}

func (m *GumballMachine) EjectQuarter() {
	switch m.state {
	case HasQuarter:
		println("Quarter returned")
		m.state = NoQuarter
	case NoQuarter:
		println("You haven't inserted a quarter")
	case SoldOut:
		println("You can't eject, you haven't inserted a quarter yet")
	case Sold:
		println("Sorry, you already turned the crank")
	}
}

func (m *GumballMachine) TurnCrank() {
	switch m.state {
	case Sold:
		println("Turning twice doesn't get you another gumball!")
	case NoQuarter:
		println("You turned but there's no quarter")
	case SoldOut:
		println("You turned, but there are no gumballs")
	case HasQuarter:
		println("You turned...")
		m.state = Sold
		m.dispense()
	}
}

func (m *GumballMachine) dispense() {
	switch m.state {
	case Sold:
		println("A gumball comes rolling out the slot")
		m.count--
		if m.count == 0 {
			println("Oops, out of gumballs!")
			m.state = SoldOut
		} else {
			m.state = NoQuarter
		}
	case NoQuarter:
		println("You need to pay first")
	case SoldOut:
		println("No gumball dispensed")
	case HasQuarter:
		println("No gumball dispensed")
	}
}

func (m *GumballMachine) String() string {
	return "Mighty Gumball, Inc.\n" +
		"Gooball Machine Model #2004\n" +
		"Inventory: " + strconv.Itoa(m.count) + " gumballs\n" +
		"Machine is " + m.stateString()
}

func (m *GumballMachine) stateString() string {
	switch m.state {
	case Sold:
		return "delivering a gumball"
	case NoQuarter:
		return "waiting for quarter"
	case HasQuarter:
		return "waiting for turn of crank"
	case SoldOut:
		return "sold out"
	default:
		return "unknown state"
	}
}

func (m *GumballMachine) Refill(count int) {
	m.count = count
	m.state = NoQuarter
}

func Simulate() {
	m := NewGumballMachine(5)

	fmt.Println(m)

	m.InsertQuarter()
	m.TurnCrank()

	fmt.Println(m)

	m.InsertQuarter()
	m.EjectQuarter()
	m.TurnCrank()

	fmt.Println(m)

	m.InsertQuarter()
	m.TurnCrank()
	m.InsertQuarter()
	m.TurnCrank()
	m.EjectQuarter()

	fmt.Println(m)

	m.InsertQuarter()
	m.InsertQuarter()
	m.TurnCrank()
	m.InsertQuarter()
	m.TurnCrank()
	m.InsertQuarter()
	m.TurnCrank()

	fmt.Println(m)
}
