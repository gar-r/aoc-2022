package main

import "fmt"

type Crt struct {
	Pixels []string
}

func (c *Crt) Draw(cycle int, regx int) {
	pos := (cycle - 1) % 40
	if regx-1 <= pos && pos <= regx+1 {
		c.Pixels[cycle-1] = "#"
	} else {
		c.Pixels[cycle-1] = "."
	}
}

func (c *Crt) Print() {
	for i := 0; i < 6; i++ {
		offset := i * 40
		for j := 0; j < 40; j++ {
			fmt.Print(c.Pixels[offset+j])
		}
		fmt.Println()
	}
}
