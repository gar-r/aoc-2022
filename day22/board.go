package main

type Board [][]Tile

type Tile int

const (
	Empty Tile = iota
	Free
	Wall
)

func (b Board) Next(x, y int, d Direction) (nx, ny int) {
	dx, dy := d.Offset()
	i := 1
	for {
		nx = (x + i*dx) % len(b)
		ny = (y + i*dy) % len(b[0])
		if nx < 0 {
			nx = len(b) + nx
		}
		if ny < 0 {
			ny = len(b[0]) + ny
		}
		if b[nx][ny] != Empty {
			return
		}
		i++
	}
}
