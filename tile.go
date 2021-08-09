package main

import (
	"fmt"
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

func CreateRandomTileMap() TileMap {
	size := 10

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	tMap := createEmptyTileMap(size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			tMap[i][j] = Tile{
				Weight: r.Intn(3),
			}
		}

	}

	return tMap
}

func (tile Tile) render() string {
	switch tile.Weight {
	case 0:
		return ","
	case 1:
		return ";"
	case 2:
		return "!"
	default:
		return "."
	}
}

func (tilesMap TileMap) Render() {
	str := ""
	for _, tiles := range tilesMap {
		for _, tile := range tiles {
			str += tile.render()
		}
		str += "\n"
	}
	fmt.Println(str)
}
