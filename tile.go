package main

import (
	"fmt"
	"github.com/Pyroarsonist/map-gonerator/helpers"
	"github.com/fatih/color"
	"math/rand"
	"strings"
)

type Tile struct {
	weight              int
	threatLevel         int
	heroPresence        bool
	destinationPresence bool
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
				neighbours = append(neighbours, tMap[i-1][j].weight)
			}
			/**
			. n .
			. x
			*/
			if j > 0 {
				neighbours = append(neighbours, tMap[i][j-1].weight)
			}
			/**
			n . .
			. x
			*/
			if i > 0 && j > 0 {
				neighbours = append(neighbours, tMap[i-1][j-1].weight)
			}
			/**
			. . n
			. x
			*/
			if i < config.size-1 && j > 0 {
				neighbours = append(neighbours, tMap[i+1][j-1].weight)
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

			helpers.MakeUnique(&weightArr)

			weight := helpers.GetRandomItem(weightArr)

			tMap[i][j] = Tile{
				weight:      weight,
				threatLevel: rand.Intn(config.maxThreatLevel),
			}
		}

	}

	return tMap
}

func getColours() (func(a ...interface{}) string, func(a ...interface{}) string, func(a ...interface{}) string, func(a ...interface{}) string, func(a ...interface{}) string) {
	white := color.New(color.FgWhite, color.Faint).SprintFunc()
	cyan := color.New(color.FgHiCyan, color.ReverseVideo, color.Bold).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	return white, cyan, green, yellow, red
}

func (tile Tile) renderTopology() string {
	_, cyan, green, yellow, red := getColours()
	if tile.heroPresence {
		return cyan("X")
	}
	if tile.destinationPresence {
		return cyan("D")
	}
	defaultMinSymbol := green(".")
	defaultMaxSymbol := red("$")
	//todo: add water tiles (negative values)
	asciiSymbols := [...]string{",", ";", "!", "v", "l", "L", "F", "E"}
	if tile.weight < 0 {
		return defaultMinSymbol
	}

	if tile.weight > len(asciiSymbols)-1 {
		return defaultMaxSymbol
	}

	colorFunc := green
	if tile.weight > 2 {
		colorFunc = yellow
	}
	if tile.weight > 4 {
		colorFunc = red
	}

	return colorFunc(asciiSymbols[tile.weight])
}

func (tile Tile) renderThreat() string {
	_, cyan, green, yellow, red := getColours()
	if tile.heroPresence {
		return cyan("X")
	}
	if tile.destinationPresence {
		return cyan("D")
	}
	defaultMinSymbol := green(".")
	defaultMaxSymbol := red("!")
	asciiSymbols := [...]string{".", ";", "!"}
	if tile.threatLevel < 0 {
		return defaultMinSymbol
	}

	if tile.threatLevel > len(asciiSymbols)-1 {
		return defaultMaxSymbol
	}

	colorFunc := green
	if tile.weight > 0 {
		colorFunc = yellow
	}
	if tile.weight > 1 {
		colorFunc = red
	}

	return colorFunc(asciiSymbols[tile.threatLevel])
}

func (tileMap TileMap) RenderedMap() (mr MapRender) {
	white, _, _, _, _ := getColours()
	for _, tileRow := range tileMap {
		mr.topology += white(strings.Repeat(".-", len(tileRow)) + ".\n")
		mr.threat += white(strings.Repeat(".-", len(tileRow)) + ".\n")

		for _, tile := range tileRow {
			mr.topology += white("|")
			mr.threat += white("|")
			mr.topology += tile.renderTopology()
			mr.threat += tile.renderThreat()
		}
		mr.topology += white("|\n")
		mr.threat += white("|\n")
	}
	mr.topology += white(strings.Repeat(".-", len(tileMap)) + ".\n")
	mr.threat += white(strings.Repeat(".-", len(tileMap)) + ".\n")
	return mr
}

func (mr MapRender) convertRenderedMapToString() string {
	return fmt.Sprintf("Topology:\n%s\nThreat:\n%s", mr.topology, mr.threat)
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
