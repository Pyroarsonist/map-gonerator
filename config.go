package main

import (
	"flag"
)

type Config struct {
	size           int
	fileOutput     string
	maxTileWeight  int
	maxHeightDiff  int
	maxThreatLevel int
	maxHP          int
}

func LoadConfig() (c Config) {
	flag.IntVar(&c.size, "size", 10, "Map size")
	flag.IntVar(&c.maxTileWeight, "mtw", 8, "Maximum tile weight")
	flag.IntVar(&c.maxHeightDiff, "mhf", 3, "Maximum height difference between neighboring tiles")
	flag.IntVar(&c.maxThreatLevel, "t", 3, "Maximum threat level")
	flag.IntVar(&c.maxHP, "hp", 15, "Maximum hp")
	flag.StringVar(&c.fileOutput, "out", "", "Should save to file")
	flag.Parse()

	return
}
