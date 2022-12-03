package main

import (
	"bufio"
	"os"
)

func readInput() []string {
	f, err := os.Open("input.dat")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	result := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}
