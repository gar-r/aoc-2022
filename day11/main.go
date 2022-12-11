package main

import (
	"fmt"
	"sort"
)

func main() {
	stats := simulate(input, 20)
	inspects := stats.InspectValues()
	sort.Sort(sort.Reverse(sort.IntSlice(inspects)))
	fmt.Printf("first top inspect count: %d\n", inspects[0])
	fmt.Printf("second top inspect count: %d\n", inspects[1])
	fmt.Printf("monkey score: %d\n", inspects[0]*inspects[1])
}

func simulate(monkeys []*Monkey, rounds int) *Stats {
	stats := NewStats()
	for i := 0; i < rounds; i++ {
		stats.Merge(simulateRound(monkeys))
	}
	return stats
}

func simulateRound(monkeys []*Monkey) *Stats {
	s := NewStats()
	for _, m := range monkeys {
		for range m.Items {
			m.InspectItem()
			m.FinishInspect()
			s.Record(m)
			m.ThrowItem(monkeys)
		}
	}
	return s
}
