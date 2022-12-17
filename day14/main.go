package main

import (
	"day14/cave"
	"fmt"
)

func main() {
	sim1()
	sim2()
}

func sim2() {
	c := cave.Parse(data, 1000, 175)
	source := &cave.Pos{W: 500}
	c.Set(source, cave.Source)
	d := c.ActualDepth() + 2
	for w := 0; w < 1000; w++ {
		c.Set(&cave.Pos{W: w, D: d}, cave.Brick)
	}
	fmt.Printf("units of sand: %d\n", c.RunSim2(source))
}

func sim1() {
	c := cave.Parse(data, 1000, 175)
	source := &cave.Pos{W: 500}
	c.Set(source, cave.Source)
	fmt.Printf("units of sand: %d\n", c.RunSim(source))
}
