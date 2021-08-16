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
	}
}

func simulateTrip(tileMap TileMap, config Config) {
	fmt.Println("Simulating trip...")

	trip := tileMap.createRandomTrip(config.maxHP)

	trip.renderTrip(tileMap)
}

func (trip Trip) cleanupTrip(tileMap TileMap) {
	tileMap[trip.hero.coordinate.height][trip.hero.coordinate.width].setHeroPresence()
	tileMap[trip.destination.height][trip.destination.width].setDestinationPresence()

}

func (trip Trip) renderTrip(tileMap TileMap) {
	trip.cleanupTrip(tileMap)
	renderedMapString := tileMap.RenderedMap().convertRenderedMapToString()
	printToConsole(renderedMapString)

}
