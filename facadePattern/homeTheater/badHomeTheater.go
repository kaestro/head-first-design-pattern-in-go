package homeTheater

import (
	"fmt"
)

type PopcornPopper struct {
}

func (p *PopcornPopper) On() {
	fmt.Println("Popcorn Popper on")
}

func (p *PopcornPopper) Pop() {
	fmt.Println("Popcorn Popper popping popcorn")
}

func (p *PopcornPopper) Off() {
	fmt.Println("Popcorn Popper off")
}

type TheaterLights struct {
}

func (t *TheaterLights) On() {
	fmt.Println("Theater Lights on")
}

func (t *TheaterLights) Dim(level int) {
	fmt.Println("Theater Lights dimming to", level, "%")
}

func (t *TheaterLights) Off() {
	fmt.Println("Theater Lights off")
}

type Screen struct {
}

func (s *Screen) Up() {
	fmt.Println("Screen up")
}

func (s *Screen) Down() {
	fmt.Println("Screen down")
}

type Projector struct {
}

func (p *Projector) On() {
	fmt.Println("Projector on")
}

func (p *Projector) Off() {
	fmt.Println("Projector off")
}

func (p *Projector) SetInput(input string) {
	fmt.Println("Projector input set to", input)
}

func (p *Projector) WideScreenMode() {
	fmt.Println("Projector wide screen mode")
}

type Player struct {
}

func (p *Player) On() {
	fmt.Println("Player on")
}

func (p *Player) Off() {
	fmt.Println("Player off")
}

func (p *Player) Play(movie string) {
	fmt.Println("Player playing", movie)
}

func SimulateBadHomeTheater() {
	popcornPopper := PopcornPopper{}
	theaterLights := TheaterLights{}
	screen := Screen{}
	projector := Projector{}
	player := Player{}

	popcornPopper.On()
	popcornPopper.Off()
	theaterLights.On()
	screen.Down()
	projector.On()
	projector.SetInput("DVD")
	projector.WideScreenMode()
	player.On()
	player.Play("The Matrix")
	player.Off()
	projector.Off()
	screen.Up()
	theaterLights.Off()
}
