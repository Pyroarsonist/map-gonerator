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
	renderedMap := tMap.RenderedMap()

	if config.fileOutput != "" {
		printToFile(renderedMap, config.fileOutput)
	} else {
		printToConsole(renderedMap)
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
