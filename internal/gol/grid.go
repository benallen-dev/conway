package gol

import (
	"math/rand"
)

const (
	alive = ""
	dead = " "
	cursor = ""
)

type Grid struct {
	Width int
	Height int
	Cells [][]Cell
}

func NewGrid(width, height int) *Grid {
	cells := make([][]Cell, height)
	for i := range cells {
		cells[i] = make([]Cell, width)
	}
	return &Grid{Width: width, Height: height, Cells: cells}
}

func (g *Grid) GetCell(x, y int) *Cell {
	return &g.Cells[y][x]
}

func (g *Grid) SetCell(x, y int, cell Cell) {
	g.Cells[y][x] = cell
}

func (g* Grid) String() string {
	var out string

	for y := range g.Height {
		for x := range g.Width {
			if (g.GetCell(x,y).IsAlive()) {
				out += alive
			} else {
				out += dead
			}
		}

		out += "\n"
	}

	return out
}


func (g *Grid) Tick(tickrate int) {
	for range tickrate {
		x := rand.Intn(g.Width)
		y := rand.Intn(g.Height)

		cell := g.GetCell(x, y)
		cell.Alive = !cell.Alive
	}
}
