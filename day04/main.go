package main

import "fmt"

func main() {
	crews := parseCrews()
	redundant := 0
	overlapping := 0
	for _, crew := range crews {
		if crew.isRedundant() {
			redundant++
			overlapping++
		} else if crew.hasOverlaps() {
			overlapping++
		}
	}
	fmt.Printf("redundant zones: %d\n", redundant)
	fmt.Printf("overlapping zones: %d\n", overlapping)
}

func parseCrews() []*crew {
	crews := make([]*crew, 0)
	data := readInput()
	for _, line := range data {
		crews = append(crews, newCrew(line))
	}
	return crews
}
