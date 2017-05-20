package main

import (
	"github.com/spaiz/gogameoflife"
	"flag"
	"fmt"
	"log"
)

var flagInputFile string
var flagDelay int
var flagDensity int
var flagRandom int

func init() {
	flag.IntVar(&flagRandom, "random", 0, "You can run game wirh random input data. Provide random value as positive integer (defines dec size)")
	flag.IntVar(&flagDensity, "density", 2, "Density of initial population in random mode. Use any integer. Less number = denser population")
	flag.IntVar(&flagDelay, "delay", 500, "You can provide delay to be for colonians generation. Negative or zero delay means no delay.")

	initExampleData := `You may provide file path, that will be used as initial data. If not, we will try to search for default
	filename in current directory.
	Data in file must be aligned NxN where N is number of symbols in each line should equal lines number.
	For exemple (capital O used):

	---O---
	---O---
	---O---

	`

	flag.StringVar(&flagInputFile, "file", "./input.txt", fmt.Sprintf(initExampleData))
	flag.Parse()
}

// Game of Life naive implementation
// For running, input data provider must be passed
// Yo can change input data source by implementing your own data providers
func main() {
	var provider gogameoflife.DataSourceProvider

	if flagRandom > 0 {
		provider = gogameoflife.NewRandomSourceDataProvider(flagRandom, flagDensity)
	} else {
		provider = gogameoflife.NewFileSourceDataProvider(flagInputFile)
	}

	game := gogameoflife.NewGame(provider)
	err := game.Start(flagDelay)

	if err != nil {
		log.Fatalf("Sorry, some error occurred: %s", err.Error())
	}
}