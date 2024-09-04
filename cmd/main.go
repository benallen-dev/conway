package main

import (
	"fmt"
	"time"

	"golang.org/x/term"

	"github.com/benallen-dev/conway/internal/gol"
)

func initDisplay() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H\033[2J")

	// hide cursor
	fmt.Print("\033[?25l")
}

func restoreDisplay() {
	// show cursor
	fmt.Print("\033[?25h")
}

func display(grid *gol.Grid, count int) {
	fmt.Printf("\033[0;0H")
	fmt.Print(grid)
	fmt.Printf("Tick: %d\n", count)
}

var (
	interval = 516 // how often to update the display in milliseconds
	tickrate = 1 // How many grid updates to do before displaying
)

func main() {
	width, height, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}

	// scaling
	width = width - 2
	height = height - 5

	grid := gol.NewGrid(width, height)

	// inital conditions
	grid.SetCell(2, 1, true)
	grid.SetCell(2, 2, true)
	grid.SetCell(2, 3, true)

	initDisplay()

	count := 0

	// TODO: Use goroutine and channel to separate display from simulation
	for {
		count = count + 1
		display(grid, count)

		grid = grid.Tick(tickrate)
	
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
