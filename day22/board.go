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

func (b Board) Wrap(x, y int, d Direction) (nx, ny int, nd Direction) {
	if x == 199 && between(y, 0, 50) && d == Down {
		return 0, y + 100, Down // edge U --> R
	}
	if between(x, 150, 200) && y == 49 && d == Right {
		return 149, x - 100, Up // edge U --> F
	}
	if between(x, 150, 200) && y == 0 && d == Left {
		return 0, x - 100, Down // edge U --> B
	}
	if x == 100 && between(y, 0, 50) && d == Up {
		return y + 50, 50, Right // edge L --> D
	}
	if between(x, 100, 150) && y == 0 && d == Left {
		return 49 - (x - 100), 50, Right // edge L --> B
	}
	if x == 149 && between(y, 50, 100) && d == Down {
		return y + 100, 49, Left // edge F --> U
	}
	if between(x, 100, 150) && y == 99 && d == Right {
		return 49 - (x - 100), 149, Left // edge F --> R
	}
	if between(x, 50, 100) && y == 50 && d == Left {
		return 100, x - 50, Down // edge D --> L
	}
	if between(x, 50, 100) && y == 99 && d == Right {
		return 49, x + 50, Up // edge D --> R
	}
	if between(x, 0, 50) && y == 50 && d == Left {
		return 149 - x, 0, Right // edge B --> L
	}
	if x == 0 && between(y, 50, 100) && d == Up {
		return 100 + y, 0, Right // edge B --> U
	}
	if x == 0 && between(y, 100, 150) && d == Up {
		return 199, y - 100, Up // edge R --> U
	}
	if between(x, 0, 50) && y == 149 && d == Right {
		return 149 - x, 99, Left // edge R --> F
	}
	if x == 49 && between(y, 100, 150) && d == Down {
		return y - 50, 99, Left
	}
	dx, dy := d.Offset()
	nx = x + dx
	ny = y + dy
	return nx, ny, d
}

func between(i, lower, upper int) bool {
	return i >= lower && i < upper
}
