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

func display(grid *gol.Grid) {
	fmt.Printf("\033[0;0H")
	fmt.Print(grid)
}

var (
	interval = 16 // how often to update the display in milliseconds
	tickrate = 50 // How many grid updates to do before displaying

)

func main() {
	width, height, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}

	// scaling
	width = width - 2
	height = height - 2

	grid := gol.NewGrid(width, height)

	// inital conditions
	grid.SetCell(3, 1, gol.Cell{Alive: true})
	grid.SetCell(2, 2, gol.Cell{Alive: true})
	grid.SetCell(1, 3, gol.Cell{Alive: true})

	initDisplay()


	// TODO: Use goroutine and channel to separate display from simulation
	for {
		display(grid)

		grid.Tick(tickrate)
	
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
