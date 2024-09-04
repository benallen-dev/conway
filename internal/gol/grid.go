package gol

import (
	"math/rand"
)

const (
	alive  = ""
	dead   = " "
	cursor = ""
)

type Grid struct {
	Width  int
	Height int
	Cells  [][]Cell // 2D slice of cells, indexed by [x][y] (width, height)
	Cursor [2]int   // 2-element array of x, y coordinates
}

func NewGrid(width, height int) *Grid {
	cells := make([][]Cell, width)
	for i := range cells {
		cells[i] = make([]Cell, height)
	}
	return &Grid{Width: width, Height: height, Cells: cells}
}

func (g *Grid) Cell(x, y int) *Cell {
	return &g.Cells[x][y]
}

func (g *Grid) SetCell(x, y int, alive bool) {
	g.Cells[x][y].alive = alive
}

func (g *Grid) String() string {
	var out string

	for y := range g.Height {
		for x := range g.Width {
			if g.Cell(x, y).IsAlive() {
				out += alive
			} else {
				out += dead
			}
		}

		out += "\n"
	}

	return out
}

func (g *Grid) CountAliveNeighbours(x, y int) int {
	count := 0

	for yy := y - 1; yy <= y+1; yy = yy + 1 {
		for xx := x - 1; xx <= x+1; xx = xx + 1 {
			if (x == xx && y == yy) || xx < 0 || yy < 0 || xx >= g.Width || yy >= g.Height {
				continue
			}

			if g.Cell(xx, yy).IsAlive() {
				count = count + 1
			}

		}
	}

	return count
}

func (g *Grid) Next() *Grid {
	newGrid := NewGrid(g.Width, g.Height)

	for x := range g.Width {
		for y := range g.Height {
			count := g.CountAliveNeighbours(x, y)

			me := g.Cell(x, y)

			// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
			if me.alive && count < 2 {
				newGrid.SetCell(x, y, false)
			}

			// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
			if !me.alive && count == 3 {
				newGrid.SetCell(x, y, true)
			}

			// Any live cell with more than three live neighbours dies, as if by overpopulation.
			if me.alive && count > 3 {
				newGrid.SetCell(x, y, false)
			}

			// Any live cell with two or three live neighbours lives on to the next generation.
			if me.alive && (count == 2 || count == 3) {
				newGrid.SetCell(x, y, true)
			}
		}
	}
	return newGrid
}

func (g *Grid) Tick(tickrate int) *Grid {
	newGrid := g.Next()
	for range tickrate - 1 {
		newGrid = newGrid.Next()
	}

	return newGrid
}

func (g *Grid) RandomChange(iterations int) {
	for range iterations {
		x := rand.Intn(g.Width)
		y := rand.Intn(g.Height)

		cell := g.Cell(x, y)
		cell.alive = !cell.alive
	}
}
