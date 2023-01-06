package main

type Path []Instruction

type Instruction interface{}

type Move struct {
	Amount int
}

type Turn struct {
	Dir Direction
}

type Direction int

const (
	Right Direction = iota
	Down
	Left
	Up
)

func (d Direction) Turn(td Direction) Direction {
	switch td {
	case Left:
		switch d {
		case Up:
			return Left
		case Right:
			return Up
		case Down:
			return Right
		default:
			return Down
		}
	case Right:
		switch d {
		case Up:
			return Right
		case Right:
			return Down
		case Down:
			return Left
		default:
			return Up
		}
	default:
		return d
	}
}

func (d Direction) Offset() (dx, dy int) {
	switch d {
	case Left:
		return 0, -1
	case Right:
		return 0, 1
	case Up:
		return -1, 0
	default:
		return 1, 0
	}
}

func (d Direction) String() string {
	switch d {
	case Left:
		return "Left"
	case Right:
		return "Right"
	case Up:
		return "Up"
	case Down:
		return "Down"
	default:
		panic(d)
	}
}
