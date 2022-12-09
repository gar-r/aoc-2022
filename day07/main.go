package main

import (
	"fmt"
)

func main() {
	commands := ParseCommands(input)
	root := traceCommands(commands)

	part1(root)
	part2(root)
}

func part1(root *Directory) {
	threshold := 100000
	total := countBelowThreshold(root, threshold)
	fmt.Printf("total: %d\n", total)
}

func countBelowThreshold(d *Directory, threshold int) int {
	total := 0
	if d.Size() < threshold {
		total += d.Size()
	}
	for _, c := range d.GetDirectories() {
		total += countBelowThreshold(c, threshold)
	}
	return total
}

func part2(root *Directory) {
	disk := 70000000
	used := root.Size()
	free := disk - used

	target := 30000000
	cleanup := target - free

	candidates := collectAboveThreshold(root, cleanup)
	optimal := min(candidates)
	fmt.Printf("optimal delete: %d\n", optimal)
}

func min(arr []int) int {
	r := arr[0]
	for i := 1; i < len(arr); i++ {
		item := arr[i]
		if item < r {
			r = item
		}
	}
	return r
}

func collectAboveThreshold(d *Directory, threshold int) []int {
	result := make([]int, 0)
	if d.Size() > threshold {
		result = append(result, d.Size())
	}
	for _, c := range d.GetDirectories() {
		result = append(result, collectAboveThreshold(c, threshold)...)
	}
	return result
}
