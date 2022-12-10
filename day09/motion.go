package main

import (
	"strconv"
	"strings"
)

type motion struct {
	dir   direction
	steps int
}

type direction struct {
	dx, dy int
}

func (d *direction) null() bool {
	return d.dx == 0 && d.dy == 0
}

func parseMotion(s string) *motion {
	parts := strings.Split(s, " ")
	n, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	return &motion{
		dir:   parseDirection(parts[0]),
		steps: n,
	}
}

func parseDirection(s string) direction {
	switch s {
	case "U":
		return direction{0, 1}
	case "D":
		return direction{0, -1}
	case "L":
		return direction{-1, 0}
	case "R":
		return direction{1, 0}
	default:
		panic(s)
	}
}
