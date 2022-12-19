package main

import (
	"fmt"
	"math/big"
)

func main() {
	part1(Input)
	part2(Input)
}

func part2(m *Map) {
	max := 4000000
	var p *Point
Search:
	for y := 0; y <= max; y++ {
		percent := float64(y) / float64(max) * 100
		fmt.Printf("%f%%\n", percent)
		uncovered := m.FindUncovered(y)
		for _, v := range uncovered {
			for _, n := range v.Enumerate() {
				if n >= 0 && n <= max {
					p = NewPoint(n, y)
					break Search
				}
			}
		}
	}
	fmt.Printf("beacon coordinates: %s\n", p)
	freq := big.NewInt(4000000)
	freq.Mul(freq, big.NewInt(int64(p.X)))
	freq.Add(freq, big.NewInt(int64(p.Y)))
	fmt.Printf("tuning frequency: %s\n", freq)
}

func part1(m *Map) {
	y := 2000000
	v := m.FullRange()
	uncovered := m.FindUncovered(y)
	for _, i := range uncovered {
		v = SubtractAll(v, i)
	}
	count := 0
	for _, i := range v {
		for _, n := range i.Enumerate() {
			p := NewPoint(n, y)
			if m.IsSensor(p) || m.IsBeacon(p) {
				continue
			}
			count++
		}
	}
	fmt.Printf("covered tile count: %d\n", count)
}
