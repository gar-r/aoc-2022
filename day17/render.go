package main

import (
	"fmt"
	"os"
)

func (m *Map) render(rock Rock, x, y int) {
	stop := len(m.data) - 25
	if stop < 0 {
		stop = 0
	}
	for i := len(m.data) - 1; i >= stop; i-- {
		row := m.data[i]
		fmt.Printf("[%2d] ", i)
		for j := 0; j < len(row); j++ {
			b := false
			for _, block := range rock.BlockPattern() {
				nx, ny := block.Apply(x, y)
				if ny == i && nx == j {
					b = true
				}
			}
			if b {
				fmt.Print("@")
			} else if m.data[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	os.Stdin.Read(make([]byte, 1))
}
