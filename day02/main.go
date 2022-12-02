package main

import "fmt"

func main() {
	input, err := readInput()
	if err != nil {
		panic(err)
	}
	score := 0
	for _, round := range input {
		score += int(round.desiredState)
		switch round.desiredState {
		case Loss:
			score += int(round.opponent.WinsTo())
		case Draw:
			score += int(round.opponent)
		case Win:
			score += int(round.opponent.LosesTo())
		}
	}
	fmt.Printf("total score: %d", score)
}
