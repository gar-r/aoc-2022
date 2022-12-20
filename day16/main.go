package main

import "fmt"

func main() {
	valves := parse(input)
	g := NewGraph(valves, "AA")
	fmt.Printf("pressure: %d\n", g.CalcPressure(30))
	fmt.Printf("pressure with elephant: %d\n", g.ElephantSim(26))
}
