package main

import "fmt"

func main() {
	input := readInput()
	part1(input)
	part2(input)
}

func part2(input []string) {
	score := 0
	for i := 0; i < len(input); i += 3 {
		g := newGroup(input[i], input[i+1], input[i+2])
		badge, ok := g.findBadge()
		if !ok {
			panic(fmt.Sprintf("badge not found: %v", g))
		}
		score += badge.priority()
	}
	fmt.Printf("group total score: %d\n", score)

}

func part1(input []string) {
	rucksacks := make([]*rucksack, len(input))
	for i, line := range input {
		rucksacks[i] = newRucksack(line)
	}

	score := 0
	for _, r := range rucksacks {
		dup, ok := r.findDuplicate()
		if !ok {
			panic(fmt.Sprintf("no duplicate in: %v", r))
		}
		score += dup.priority()
	}

	fmt.Printf("total score: %d\n", score)
}
