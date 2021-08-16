package main

import "fmt"

type AlgorithmName int

const (
	Greedy AlgorithmName = iota
)

func (trip *Trip) simulateWithAlgorithm(algoName AlgorithmName) {
	destination := trip.destination
	hero := &trip.hero
	for {
		trip.day += 1

		same, nextCoordinate := getNextCoordinate(algoName, destination, hero.coordinate)
		if same {
			return
		}

		hero.coordinate = nextCoordinate
		newTile := &trip.tileMap[nextCoordinate.height][nextCoordinate.width]
		hero.hpLoss(newTile.threatLevel)
		newTile.setHeroTrace()
		if hero.hp == 0 {
			newTile.setHeroDeath()
			return
		}
	}
}

func getNextCoordinate(algoName AlgorithmName, destination TileCoordinate, heroC TileCoordinate) (same bool, nextCoordinate TileCoordinate) {
	switch algoName {
	case Greedy:
		return getNextCoordinateGreedy(destination, heroC)
	default:
		panic(fmt.Sprintln("Invalid algorithm name:", algoName))
	}

}

func getNextCoordinateGreedy(destination TileCoordinate, heroC TileCoordinate) (same bool, nextCoordinate TileCoordinate) {
	nextCoordinate = heroC

	if destination.height == heroC.height {
		if destination.width == heroC.width {
			return true, nextCoordinate
		}

		nextCoordinate.width = heroC.width + 1
		if destination.width < heroC.width {
			nextCoordinate.width = heroC.width - 1
		}
		return false, nextCoordinate

	}

	nextCoordinate.height = heroC.height + 1
	if destination.height < heroC.height {
		nextCoordinate.height = heroC.height - 1
	}
	return false, nextCoordinate
}
