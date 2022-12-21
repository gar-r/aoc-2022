package main

import (
	"fmt"
)

func main() {
	part1()
	part2()
}

func spawnRocks(m *Map, count int) int {
	for i := 0; i < count; i++ {
		m.SimRock()
	}
	_, y := m.Peak()
	return y
}

func part1() {
	m := NewMap(input)
	h := spawnRocks(m, 2022)
	fmt.Printf("%d\n", h+1)
}

func part2() {
	s := input
	start := 430
	flen := 1700
	dumpPattern(s) // frequency can be determined from dump
	total := 1000000000000
	m := NewMap(s)
	h1 := spawnRocks(m, start-1)
	h2 := spawnRocks(m, flen)
	dh := h2 - h1
	repeats := (total - start) / flen
	chunk := repeats * dh
	rem := total - (start + repeats*flen) + 1
	m = NewMap(s)
	h1 = spawnRocks(m, start-1)
	h2 = spawnRocks(m, rem)
	extra := h2 - h1
	h := h1 + chunk + extra
	fmt.Printf("%d\n", h+1)
}

func dumpPattern(s string) {
	m := NewMap(s)
	prev := 0
	for i := 0; i < 10000; i++ {
		m.SimRock()
		_, y := m.Peak()
		fmt.Printf("%1d", y-prev)
		prev = y
	}
	fmt.Println()
}
