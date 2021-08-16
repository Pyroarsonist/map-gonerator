package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	config := LoadConfig()

	if config.version {
		viewVersion()
		return
	}

	tileMap := CreateRandomTileMap(config)
	renderedMapString := tileMap.RenderedMap().convertRenderedMapToString()

	if config.fileOutput != "" {
		printToFile(renderedMapString, config.fileOutput)
	} else {
		printToConsole(renderedMapString)
		simulateTrip(tileMap, config)
	}
}

func viewVersion() {
	version, err := ioutil.ReadFile("version.txt")
	if err != nil {
		panic("failed opening version file")
	}
	fmt.Printf("Version %s\n", version)
}
