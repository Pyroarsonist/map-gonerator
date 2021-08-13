package main

import (
	"math/rand"
	"time"
)

type Tile struct {
	Weight int
}

type Tiles []Tile

type TileMap []Tiles

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
				weightArr = append(weightArr, r.Intn(config.maxTileWeight))
			}

			makeUnique(&weightArr)

			weight := getRandom(weightArr)

			tMap[i][j] = Tile{
				Weight: weight,
			}
		}

	}

	return tMap
}

func (tile Tile) render() string {
	defaultSymbol := "."
	asciiSymbols := [...]string{",", ";", "!", "v", "l", "L", "F", "E", "$"}
	if tile.Weight < 0 || tile.Weight > len(asciiSymbols)-1 {
		return defaultSymbol
	}

	return asciiSymbols[tile.Weight]
}

func (tilesMap TileMap) RenderedMap() string {
	str := ""
	for _, tiles := range tilesMap {
		for _, tile := range tiles {
			str += tile.render()
		}
		str += "\n"
	}
	return str
}
