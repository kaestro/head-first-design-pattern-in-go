package goodGumballMachine

import (
	"fmt"
	"strconv"
)

type State interface {
	InsertQuarter()
	EjectQuarter()
	TurnCrank()
	Dispense()
}

type HasQuarterState struct {
	machine *GumballMachine
}

func (s *HasQuarterState) InsertQuarter() {
	fmt.Println("You can't insert another quarter")
}

func (s *HasQuarterState) EjectQuarter() {
	fmt.Println("Quarter returned")
	s.machine.State = s.machine.NoQuarterState
}

func (s *HasQuarterState) TurnCrank() {
	fmt.Println("You turned...")
	s.machine.State = s.machine.SoldState
}

func (s *HasQuarterState) Dispense() {
	fmt.Println("No gumball dispensed")
}

type NoQuarterState struct {
	machine *GumballMachine
}

func (s *NoQuarterState) InsertQuarter() {
	fmt.Println("You inserted a quarter")
	s.machine.State = s.machine.HasQuarterState
}

func (s *NoQuarterState) EjectQuarter() {
	fmt.Println("You haven't inserted a quarter")
}

func (s *NoQuarterState) TurnCrank() {
	fmt.Println("You turned, but there's no quarter")
}

func (s *NoQuarterState) Dispense() {
	fmt.Println("You need to pay first")
}

type SoldState struct {
	machine *GumballMachine
}

func (s *SoldState) InsertQuarter() {
	fmt.Println("Please wait, we're already giving you a gumball")
}

func (s *SoldState) EjectQuarter() {
	fmt.Println("Sorry, you already turned the crank")
}

func (s *SoldState) TurnCrank() {
	fmt.Println("Turning twice doesn't get you another gumball!")
}

func (s *SoldState) Dispense() {
	fmt.Println("A gumball comes rolling out the slot")
	s.machine.ReleaseBall()
	if s.machine.count == 0 {
		fmt.Println("Oops, out of gumballs!")
		s.machine.State = s.machine.SoldOutState
	} else {
		s.machine.State = s.machine.NoQuarterState
	}
}

type SoldOutState struct {
}

func (s *SoldOutState) InsertQuarter() {
	fmt.Println("You can't insert a quarter, the machine is sold out")
}

func (s *SoldOutState) EjectQuarter() {
	fmt.Println("You can't eject, you haven't inserted a quarter yet")
}

func (s *SoldOutState) TurnCrank() {
	fmt.Println("You turned, but there are no gumballs")
}

func (s *SoldOutState) Dispense() {
	fmt.Println("No gumball dispensed")
}

type GumballMachine struct {
	State           State
	HasQuarterState State
	NoQuarterState  State
	SoldState       State
	SoldOutState    State
	count           int
}

func NewGumballMachine(count int) *GumballMachine {
	m := &GumballMachine{count: count}
	m.HasQuarterState = &HasQuarterState{machine: m}
	m.NoQuarterState = &NoQuarterState{machine: m}
	m.SoldState = &SoldState{machine: m}
	m.SoldOutState = &SoldOutState{}
	if m.count > 0 {
		m.State = m.NoQuarterState
	} else {
		m.State = m.SoldOutState
	}
	return m
}

func (m *GumballMachine) InsertQuarter() {
	m.State.InsertQuarter()
}

func (m *GumballMachine) EjectQuarter() {
	m.State.EjectQuarter()
}

func (m *GumballMachine) TurnCrank() {
	m.State.TurnCrank()
	m.State.Dispense()
}

func (m *GumballMachine) ReleaseBall() {
	if m.count != 0 {
		m.count--
	}
}

// Define String, stateString, Refill, and Simulate similarly...
func (m *GumballMachine) String() string {
	return "\nMighty Gumball, Inc.\nGoofy gumball model #2004\nInventory: " + strconv.Itoa(m.count) + " gumball"
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
}
