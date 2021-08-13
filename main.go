package main

import (
	"fmt"
)

func main() {
	fmt.Println("Generating random tiles")

	config := LoadConfig()

	tileMap := CreateRandomTileMap(config)
	renderedMapString := tileMap.RenderedMap().convertRenderedMapToString()

	if config.fileOutput != "" {
		printToFile(renderedMapString, config.fileOutput)
	} else {
		printToConsole(renderedMapString)
		simulateTrip(tileMap, config)
	}
}
