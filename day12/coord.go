package main

import (
	"fmt"
)

type Coord struct {
	Row int
	Col int
}

func (c *Coord) Offset(dir Direction) *Coord {
	n := &Coord{Row: c.Row, Col: c.Col}
	switch dir {
	case North:
		n.Row--
	case East:
		n.Col++
	case South:
		n.Row++
	case West:
		n.Col--
	}
	return n
}

func (c *Coord) Equal(d *Coord) bool {
	return c.Row == d.Row && c.Col == d.Col
}

func (c *Coord) String() string {
	return fmt.Sprintf("(%d,%d)", c.Row, c.Col)
}

type Direction uint

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	switch d {
	case North:
		return "^"
	case East:
		return ">"
	case South:
		return "v"
	case West:
		return "<"
	default:
		return "-"
	}
}
