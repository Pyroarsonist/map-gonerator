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

	//todo: hero and destination coords should not be the same
	return Trip{
		hero: Hero{
			coordinate: tileMap.GetRandomCoordinate(),
			hp:         r.Intn(maxHP),
		},
		destination: tileMap.GetRandomCoordinate(),
		day:         0,
	}
}

func simulateTrip(tileMap TileMap, config Config) {
	fmt.Println("Simulating trip...")

	trip := tileMap.createRandomTrip(config.maxHP)

	trip.renderTrip(tileMap)
}

func (trip Trip) cleanupTrip(tileMap TileMap) {
	tileMap[trip.hero.coordinate.height][trip.hero.coordinate.width].heroPresence = true
	tileMap[trip.destination.height][trip.destination.width].destinationPresence = true

}

func (trip Trip) renderTrip(tileMap TileMap) {
	trip.cleanupTrip(tileMap)
	renderedMapString := tileMap.RenderedMap().convertRenderedMapToString()
	printToConsole(renderedMapString)

}
