package display

import (
	"fmt"
)

func Init() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H\033[2J")

	// hide cursor
	fmt.Print("\033[?25l")
}

func Restore() {
	// show cursor
	fmt.Print("\033[?25h")
}

func Draw(grid string, count int) {
	fmt.Printf("\033[0;0H")
	fmt.Print(grid)
	fmt.Printf("\nTick: %d\n", count)
}
