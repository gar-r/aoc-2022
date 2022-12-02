package main

import "fmt"

type State int

const (
	Loss = 0
	Draw = 3
	Win  = 6
)

func ParseState(c string) (s State, err error) {
	switch c {
	case "X":
		s = Loss
	case "Y":
		s = Draw
	case "Z":
		s = Win
	default:
		err = fmt.Errorf("invalid state: %v", c)
	}
	return
}

type Figure int

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

func (f Figure) WinsTo() Figure {
	switch f {
	case Rock:
		return Scissors
	case Paper:
		return Rock
	default:
		return Paper
	}
}

func (f Figure) LosesTo() Figure {
	switch f {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	default:
		return Rock
	}
}

func ParseFigure(c string) (f Figure, err error) {
	switch c {
	case "A":
		f = Rock
	case "B":
		f = Paper
	case "C":
		f = Scissors
	default:
		err = fmt.Errorf("invalid figure: %v", c)
	}
	return
}
