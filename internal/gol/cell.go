package gol

type Cell struct {
	Alive bool
}

func (c *Cell) IsAlive() bool {
	return c.Alive
}
