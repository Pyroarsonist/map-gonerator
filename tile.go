package main

import (
	"fmt"
	"github.com/Pyroarsonist/map-gonerator/helpers"
	"math/rand"
)

type Tile struct {
	passability         int
	threatLevel         int
	isHeroStartLocation bool
	isDestination       bool
	isHeroTrace         bool
	isHeroDeath         bool
}

type Tiles []Tile

type TileMap []Tiles

type TileCoordinate struct {
	width  int
	height int
}

type MapRender struct {
	topology string
	threat   string
}

func createEmptyTileMap(size int) TileMap {
	tMap := make(TileMap, size)
	for i := 0; i < size; i++ {
		tMap[i] = make(Tiles, size)
	}

	return tMap
}

func CreateRandomTileMap(config Config) TileMap {
	fmt.Println("Generating random tiles")

	r := helpers.GetRandomGenerator()

	tMap := createEmptyTileMap(config.size)

	for i := 0; i < config.size; i++ {
		for j := 0; j < config.size; j++ {
			var neighbours []int
			/**
			. . .
			n x
			*/
			if i > 0 {
				neighbours = append(neighbours, tMap[i-1][j].passability)
			}
			/**
			. n .
			. x
			*/
			if j > 0 {
				neighbours = append(neighbours, tMap[i][j-1].passability)
			}
			/**
			n . .
			. x
			*/
			if i > 0 && j > 0 {
				neighbours = append(neighbours, tMap[i-1][j-1].passability)
			}
			/**
			. . n
			. x
			*/
			if i < config.size-1 && j > 0 {
				neighbours = append(neighbours, tMap[i+1][j-1].passability)
			}

			var passabilityArr []int
			for n := range neighbours {
				for w := n - config.maxHeightDiff; w <= n+config.maxHeightDiff; w++ {
					if w > 0 {
						passabilityArr = append(passabilityArr, w)
					}
				}
			}
			if len(passabilityArr) == 0 {
				//todo: add water tiles (negative values)
				passabilityArr = append(passabilityArr, r.Intn(config.maxTilePassability))
			}

			helpers.MakeUnique(&passabilityArr)

			passability := helpers.GetRandomItem(passabilityArr)

			tMap[i][j] = Tile{
				passability: passability,
				threatLevel: rand.Intn(config.maxThreatLevel),
			}
		}

	}

	return tMap
}

func (tileMap TileMap) GetRandomCoordinate() TileCoordinate {
	r := helpers.GetRandomGenerator()
	//todo: maybe refactor tilemap to data and size
	size := len(tileMap[0])
	return TileCoordinate{
		width:  r.Intn(size),
		height: r.Intn(size),
	}
}

func (tile *Tile) setHeroStartLocation() {
	tile.isHeroStartLocation = true
	tile.threatLevel = 0
}

func (tile *Tile) setDestinationPresence() {
	tile.isDestination = true
	tile.threatLevel = 0
}

func (tile *Tile) setHeroTrace() {
	tile.isHeroTrace = true
	tile.threatLevel = 0
}

func (tile *Tile) setHeroDeath() {
	tile.isHeroDeath = true
}
