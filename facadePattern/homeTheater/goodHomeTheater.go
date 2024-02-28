package homeTheater

type HomeTheaterFacade struct {
	player        *Player
	popcornPopper *PopcornPopper
	theaterLights *TheaterLights
	screen        *Screen
	projector     *Projector
}

func NewHomeTheaterFacade(player *Player, popcornPopper *PopcornPopper, theaterLights *TheaterLights, screen *Screen, projector *Projector) *HomeTheaterFacade {
	return &HomeTheaterFacade{player, popcornPopper, theaterLights, screen, projector}
}

func (h *HomeTheaterFacade) WatchMovie(movie string) {
	h.popcornPopper.On()
	h.popcornPopper.Pop()
	h.theaterLights.Dim(10)
	h.screen.Down()
	h.projector.On()
	h.projector.WideScreenMode()
	h.player.On()
	h.player.Play(movie)
}

func (h *HomeTheaterFacade) EndMovie() {
	h.popcornPopper.Off()
	h.theaterLights.On()
	h.screen.Up()
	h.projector.Off()
	h.player.Off()
}

func SimulateGoodHomeTheater() {
	popcornPopper := PopcornPopper{}
	theaterLights := TheaterLights{}
	screen := Screen{}
	projector := Projector{}
	player := Player{}

	homeTheater := NewHomeTheaterFacade(&player, &popcornPopper, &theaterLights, &screen, &projector)
	homeTheater.WatchMovie("Raiders of the Lost Ark")
	homeTheater.EndMovie()
}
