package main

import (
	"flag"
)

type Config struct {
	size          int
	fileOutput    string
	maxTileWeight int
	maxHeightDiff int
}

func LoadConfig() (c Config) {
	flag.IntVar(&c.size, "size", 10, "Map size")
	flag.IntVar(&c.maxTileWeight, "size", 9, "Maximum tile weight")
	flag.IntVar(&c.maxHeightDiff, "size", 3, "Maximum height difference between neighboring tiles")
	flag.StringVar(&c.fileOutput, "out", "", "Should save to file")
	flag.Parse()

	return
}
