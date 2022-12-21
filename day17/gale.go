package main

type Direction int

const (
	Left Direction = iota
	Right
)

func (d Direction) Apply(coord int) int {
	if d == Left {
		return coord - 1
	}
	return coord + 1
}

func (d Direction) String() string {
	if d == Left {
		return "<"
	}
	return ">"
}

type Gale struct {
	pattern []Direction
	step    int
}

func (g *Gale) Next() Direction {
	d := g.pattern[g.step]
	g.step++
	if g.step == len(g.pattern) {
		g.step = 0
	}
	return d
}
