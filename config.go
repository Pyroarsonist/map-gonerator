package main

import (
	"flag"
	"github.com/fatih/color"
)

type Config struct {
	size               int
	fileOutput         string
	maxTilePassability int
	maxHeightDiff      int
	maxThreatLevel     int
	maxHP              int
	version            bool
}

func LoadConfig() (c Config) {
	flag.IntVar(&c.size, "size", 10, "Map size")
	flag.IntVar(&c.maxTilePassability, "mtp", 8, "Maximum tile passability")
	flag.IntVar(&c.maxHeightDiff, "mhf", 3, "Maximum height difference between neighboring tiles")
	flag.IntVar(&c.maxThreatLevel, "t", 3, "Maximum threat level")
	flag.IntVar(&c.maxHP, "hp", 15, "Maximum hp")
	flag.StringVar(&c.fileOutput, "out", "", "Should save to file")
	flag.BoolVar(&c.version, "v", false, "View version")
	flag.Parse()

	if c.fileOutput != "" {
		color.NoColor = true
	}

	return
}
