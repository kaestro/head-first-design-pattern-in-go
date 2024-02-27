package main

import (
	"adapterPattern/dieBremerStadtmusikanten"
	"fmt"
)

func main() {
	fmt.Println("Adapter Pattern")

	mallardDuck := dieBremerStadtmusikanten.MallardDuck{}
	rocketDuck := dieBremerStadtmusikanten.RocketDuck{}
	wildTurkey := dieBremerStadtmusikanten.WildTurkey{}
	turkeyAdapter := dieBremerStadtmusikanten.NewTurkeyAdapter(&wildTurkey)

	// go style slice initialization
	musicianGroup := make([]dieBremerStadtmusikanten.Duck, 0, 3)
	musicianGroup = append(musicianGroup, &mallardDuck)
	musicianGroup = append(musicianGroup, &rocketDuck)
	musicianGroup = append(musicianGroup, turkeyAdapter)

	for _, musician := range musicianGroup {
		musician.Quack()
		musician.Fly()
	}
}
