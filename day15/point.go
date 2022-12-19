package main

import "fmt"

type Point struct {
	X, Y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func (p *Point) Equal(other *Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p *Point) Distance(other *Point) int {
	return abs(p.X-other.X) + abs(p.Y-other.Y)
}

func (p *Point) Offset(dx, dy int) *Point {
	return NewPoint(p.X+dx, p.Y+dy)
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
