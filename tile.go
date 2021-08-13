package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tile struct {
	Weight      int
	ThreatLevel int
}

type Tiles []Tile

type TileMap []Tiles

type MapRender struct {
	Topology string
	Threat   string
}

func createEmptyTileMap(size int) TileMap {
	tMap := make(TileMap, size)
	for i := 0; i < size; i++ {
		tMap[i] = make(Tiles, size)
	}

	return tMap
}

func CreateRandomTileMap(config Config) TileMap {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	tMap := createEmptyTileMap(config.size)

	for i := 0; i < config.size; i++ {
		for j := 0; j < config.size; j++ {
			var neighbours []int
			/**
			. . .
			n x
			*/
			if i > 0 {
				neighbours = append(neighbours, tMap[i-1][j].Weight)
			}
			/**
			. n .
			. x
			*/
			if j > 0 {
				neighbours = append(neighbours, tMap[i][j-1].Weight)
			}
			/**
			n . .
			. x
			*/
			if i > 0 && j > 0 {
				neighbours = append(neighbours, tMap[i-1][j-1].Weight)
			}
			/**
			. . n
			. x
			*/
			if i < config.size-1 && j > 0 {
				neighbours = append(neighbours, tMap[i+1][j-1].Weight)
			}

			var weightArr []int
			for n := range neighbours {
				for w := n - config.maxHeightDiff; w <= n+config.maxHeightDiff; w++ {
					if w > 0 {
						weightArr = append(weightArr, w)
					}
				}
			}
			if len(weightArr) == 0 {
				//todo: add water tiles (negative values)
				weightArr = append(weightArr, r.Intn(config.maxTileWeight))
			}

			makeUnique(&weightArr)

			weight := getRandom(weightArr)

			tMap[i][j] = Tile{
				Weight:      weight,
				ThreatLevel: rand.Intn(config.maxThreatLevel),
			}
		}

	}

	return tMap
}

func (tile Tile) renderTopology() string {
	defaultMinSymbol := "."
	defaultMaxSymbol := "$"
	//todo: add water tiles (negative values)
	asciiSymbols := [...]string{",", ";", "!", "v", "l", "L", "F", "E"}
	if tile.Weight < 0 {
		return defaultMinSymbol
	}

	if tile.Weight > len(asciiSymbols)-1 {
		return defaultMaxSymbol
	}

	return asciiSymbols[tile.Weight]
}

func (tile Tile) renderThreat() string {
	defaultMinSymbol := "."
	defaultMaxSymbol := "!"
	asciiSymbols := [...]string{".", ";", "!"}
	if tile.ThreatLevel < 0 {
		return defaultMinSymbol
	}

	if tile.ThreatLevel > len(asciiSymbols)-1 {
		return defaultMaxSymbol
	}

	return asciiSymbols[tile.ThreatLevel]
}

func (tilesMap TileMap) RenderedMap() (mr MapRender) {
	for _, tiles := range tilesMap {
		for _, tile := range tiles {
			mr.Topology += tile.renderTopology()
			mr.Threat += tile.renderThreat()
		}
		mr.Topology += "\n"
		mr.Threat += "\n"
	}
	return mr
}

func (mr MapRender) convertRenderedMapToString() string {
	return fmt.Sprintf("Topology:\n%s\nThreat:\n%s", mr.Topology, mr.Threat)
}
