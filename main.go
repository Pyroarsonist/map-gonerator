package main

import "fmt"

func main() {
	fmt.Println("Generating random tiles")

	tMap := CreateRandomTileMap()
	tMap.Render()
}
