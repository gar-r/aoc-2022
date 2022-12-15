package main

import (
	"day13/packet"
	"fmt"
	"sort"
)

func main() {
	input := parsePairs(inputData)
	fmt.Printf("sum of ordered item indices: %d\n",
		sumOrderedPacketIndices(input))

	packets := parseAll(inputData)
	div1 := packet.Parse("[[2]]")
	div2 := packet.Parse("[[6]]")
	packets = append(packets, div1, div2)
	sort.Slice(packets, func(i, j int) bool {
		return packets[i].Compare(packets[j]) == -1
	})
	var i1, i2 int
	for i, p := range packets {
		if p == div1 {
			i1 = i + 1
		}
		if p == div2 {
			i2 = i + 1
		}
	}

	fmt.Printf("decoder key: %d\n", i1*i2)
}

func sumOrderedPacketIndices(example []*packet.Pair) int {
	sum := 0
	for i := 1; i <= len(example); i++ {
		p := example[i-1]
		if p.IsOrdered() {
			sum += i
		}
	}
	return sum
}
