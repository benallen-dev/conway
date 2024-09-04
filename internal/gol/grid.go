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
	Cells  [][]Cell // 2D slice of cells
	Cursor [2]int   // 2-element array of x, y coordinates
}

func NewGrid(width, height int) *Grid {
	cells := make([][]Cell, height)
	for i := range cells {
		cells[i] = make([]Cell, width)
	}
	return &Grid{Width: width, Height: height, Cells: cells}
}

func (g *Grid) Cell(x, y int) *Cell {
	return &g.Cells[y][x]
}

func (g *Grid) SetCell(x, y int, alive bool) {
	g.Cells[y][x].alive = alive
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

func clamp(in, min, max int) int {
	if in < min {
		return min
	} else if in > max {
		return max
	} else {
		return in
	}
}

func (g *Grid) Neighbours(x, y int) *[]Cell {
	neighbours := []Cell{}

	xrange := []int{clamp(x-1, 0, g.Width), clamp(x+1, 0, g.Width)}
	yrange := []int{clamp(y-1, 0, g.Width), clamp(y+1, 0, g.Height)}

	for i := range xrange {
		for j := range yrange {
			if i == j {
				continue
			}
			
			neighbours = append(neighbours, *g.Cell(x,y))
		}
	}

	return &neighbours
}

func (g *Grid) MoveCursor(x, y int) {
	g.Cursor[0] = x
	g.Cursor[1] = y
}

func (g *Grid) CursorPosition() (int, int) {
	return g.Cursor[0], g.Cursor[1]
}

func (g *Grid) Next() *Grid{
	newGrid := NewGrid(g.Height, g.Width)

	for x := range g.Width {
		for y := range g.Height {
			neighbours := *g.Neighbours(x,y)
			count := 0

			for _, c := range neighbours {
				if c.IsAlive() {
					count += 1
				}
			}

			me := g.Cell(x, y)
		
			// Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
			if !me.alive && count == 3 {
				newGrid.SetCell(x, y, true)
			}

			// Any live cell with fewer than two live neighbours dies, as if by underpopulation.
			if me.alive && count < 2 {
				newGrid.SetCell(x, y, false)
			}
			
			// Any live cell with more than three live neighbours dies, as if by overpopulation.
			if me.alive && count > 3 {
				newGrid.SetCell(x, y, false)
			}

			// Any live cell with two or three live neighbours lives on to the next generation.
			// aka "do nothing"
		}
	}
		return newGrid
}
	
func (g *Grid) Tick(tickrate int) *Grid{
	newGrid := g.Next()
	for range tickrate -1 {
		newGrid = newGrid.Next()
	}

	return newGrid
}

func (g *Grid) RandomChange(tickrate int) {
				for range tickrate {
		x := rand.Intn(g.Width)
		y := rand.Intn(g.Height)

		cell := g.Cell(x, y)
		cell.alive = !cell.alive
	}
}
