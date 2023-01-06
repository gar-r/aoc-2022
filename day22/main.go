package main

import "fmt"

func main() {

	data := input

	player := &Player{
		row: 0,
		col: startPos(data.b),
		dir: Right,
	}

	player.Route(data.b, data.p)

	fmt.Printf(
		"player state:\n   row: %d\n   col: %d\n   dir: %s\n",
		player.row+1,
		player.col+1,
		player.dir,
	)

	fmt.Printf("code: %d\n", player.Code())
}

func startPos(b Board) int {
	y := 0
	for {
		if b[0][y] == Free {
			return y
		}
		y++
	}
}
