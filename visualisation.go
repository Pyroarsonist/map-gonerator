package main

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"strings"
)

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
	if tile.isHero {
		return cyan("X")
	}
	if tile.isDestination {
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
	if tile.isHero {
		return cyan("X")
	}
	if tile.isDestination {
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

func printToFile(str string, fileName string) {
	f, err := os.Create(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(str)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Written to", fileName)
	fmt.Println()
}

func printToConsole(str string) {
	fmt.Println()
	fmt.Println(str)
	fmt.Println()
}
