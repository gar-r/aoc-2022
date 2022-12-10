package main

import "fmt"

func main() {
	fmt.Printf("example motions: %d\n", len(example))
	fmt.Printf("input motions: %d\n", len(input))

	r := newRope(10)

	for _, m := range input {
		for i := 0; i < m.steps; i++ {
			r.drag(m.dir)
		}
	}

	fmt.Printf("distinct positions visited by tail: %d", len(r.tailLog))
}
