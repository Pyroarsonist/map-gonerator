package main

import (
	"fmt"
	"github.com/Pyroarsonist/map-gonerator/helpers"
)

type Hero struct {
	coordinate TileCoordinate
	hp         int
}

type Trip struct {
	hero        Hero
	destination TileCoordinate
	day         int
	tileMap     TileMap
}

func (tileMap TileMap) createRandomTrip(maxHP int) Trip {
	r := helpers.GetRandomGenerator()
	heroCoordinate := tileMap.GetRandomCoordinate()
	destinationCoordinate := tileMap.GetRandomCoordinate()
	for {
		if !(destinationCoordinate.width == heroCoordinate.width && destinationCoordinate.height == heroCoordinate.height) {
			break
		}
		destinationCoordinate = tileMap.GetRandomCoordinate()
	}

	return Trip{
		hero: Hero{
			coordinate: heroCoordinate,
			hp:         r.Intn(maxHP),
		},
		destination: destinationCoordinate,
		day:         0,
		tileMap:     tileMap,
	}
}

func simulateTrip(tileMap TileMap, config Config) {
	fmt.Println("Simulating trip...")

	trip := tileMap.createRandomTrip(config.maxHP)

	trip.start()
}

func (trip *Trip) cleanupTrip() {
	tileMap := trip.tileMap
	tileMap[trip.hero.coordinate.height][trip.hero.coordinate.width].setHeroStartLocation()
	tileMap[trip.destination.height][trip.destination.width].setDestinationPresence()

}

func (trip *Trip) start() {
	trip.cleanupTrip()
	renderedMapString := trip.tileMap.RenderedMap().convertRenderedMapToString()
	printToConsole(renderedMapString)

	fmt.Println("Starting trip")
	trip.simulateWithAlgorithm(Greedy)
	renderedMapString = trip.tileMap.RenderedMap().convertRenderedMapToString()
	printToConsole(renderedMapString)
	fmt.Println("Days gone:", trip.day)
	fmt.Println("HP remaining:", trip.hero.hp)
}

func (hero *Hero) hpLoss(value int) {
	hero.hp -= value
	if hero.hp < 0 {
		hero.hp = 0
	}
}
