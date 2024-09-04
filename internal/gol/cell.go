package gol

type Cell struct {
	alive bool
}

func (c *Cell) IsAlive() bool {
	return c.alive
}
