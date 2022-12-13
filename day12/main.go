package main

import (
	"fmt"
)

func main() {
	fmt.Println("example:")
	fmt.Printf("S -> E: %d\n", example.FindShortestPath())
	fmt.Printf("a -> E: %d\n", example.FindShortestPathToAny())

	fmt.Println("input:")
	fmt.Printf("S -> E: %d\n", input.FindShortestPath())
	fmt.Printf("a -> E: %d\n", input.FindShortestPathToAny())
}
