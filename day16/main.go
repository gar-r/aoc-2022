package main

import "fmt"

func main() {
	valves := parse(input)
	g := NewGraph(valves, "AA")
	max := g.CalcPressure(30)
	fmt.Printf("pressure: %d\n", max)
}
