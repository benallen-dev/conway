package main

import (
	"time"

	"golang.org/x/term"

	"github.com/benallen-dev/conway/internal/gol"
	"github.com/benallen-dev/conway/internal/display"
)

var (
	interval = 17 // Milliseconds between output updates
	tickrate = 1  // How many grid updates to do before displaying
)

func main() {
	width, height, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}

	width = width - 1 // Adds empty space to the right
	height = height - 5 // Adds empty space to the bottom

	grid := gol.NewGrid(width, height)
	grid.RandomChange(int(0.2 * float64(width) * float64(height))) // randomly flip cells 0.2 * # of cells in grid

	count := 0
	
	display.Init()
	display.Draw(grid.String(), count)
	time.Sleep(time.Duration(2) * time.Second) // Allow observer to see initial state

	// TODO: Separate display from simulation
	// TODO: Listen for SIGINT and restore display
	for {
		count = count + 1
		display.Draw(grid.String(), count)

		time.Sleep(time.Duration(interval) * time.Millisecond)

		grid = grid.Tick(tickrate)
	}
}
