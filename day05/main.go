package main

import "fmt"

func main() {
	moves := readMoves()
	cont1 := readContainer()
	cont2 := readContainer()
	execMoves9000(cont1, moves)
	execMoves9001(cont2, moves)
	printTop(cont1)
	printTop(cont2)
}

func printTop(cont *container) {
	for _, c := range cont.stacks {
		fmt.Print(c.pop())
	}
	fmt.Println()
}

func execMoves9000(cont *container, moves []*move) {
	for _, move := range moves {
		for i := 0; i < move.num; i++ {
			item := cont.stacks[move.from].pop()
			cont.stacks[move.to].push(item)
		}
	}
}

func execMoves9001(cont *container, moves []*move) {
	for _, move := range moves {
		tmp := &stack{}
		for i := 0; i < move.num; i++ {
			tmp.push(cont.stacks[move.from].pop())
		}
		for i := 0; i < move.num; i++ {
			cont.stacks[move.to].push(tmp.pop())
		}
	}
}

func readMoves() []*move {
	raw := readInput("moves.dat")
	moves := make([]*move, len(raw))
	for i, line := range raw {
		moves[i] = newMove(line)
	}
	return moves
}

func readContainer() *container {
	raw := readInput("stacks.dat")
	return newContainer(raw)
}
