package main

import (
	"fmt"
)

func main() {
	data := input
	part1(data)
	part2(data)
}

func part2(s string) {
	data := parseInput(s)
	left, right, _ := parseBinary(data["root"])
	b := data.eval(right)
	eq := data.solve(left, "humn")
	fmt.Printf("%s = %d\n", eq, b)
}

func part1(s string) {
	data := parseInput(s)
	root := data.eval("root")
	fmt.Printf("root value: %d\n", root)
}
