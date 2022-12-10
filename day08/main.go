package main

import "fmt"

func main() {
	forest := NewForest(input)
	fmt.Printf("visible trees: %d\n", forest.CountVisible())
	fmt.Printf("best scenic score: %d\n", forest.FindBestScore())
}
