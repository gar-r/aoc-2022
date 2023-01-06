package main

import "fmt"

func main() {
	part1()
	fmt.Println("---")
	part2()
}

func part2() {
	data := input // will only work for input due to cube topology

	player := &Player{
		row:  0,
		col:  50,
		dir:  Right,
		cube: true,
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

func part1() {
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
