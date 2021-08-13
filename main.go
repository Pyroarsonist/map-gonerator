package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Generating random tiles")

	config := LoadConfig()

	tMap := CreateRandomTileMap(config)
	renderedMapString := tMap.RenderedMap().convertRenderedMapToString()

	if config.fileOutput != "" {
		printToFile(renderedMapString, config.fileOutput)
	} else {
		printToConsole(renderedMapString)
	}
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
}

func printToConsole(str string) {
	fmt.Println()
	fmt.Println(str)
}
