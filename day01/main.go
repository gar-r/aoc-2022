package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := readInput()
	if err != nil {
		panic(err)
	}
	i1, max1 := findMax(data)
	delete(data, i1)
	i2, max2 := findMax(data)
	delete(data, i2)
	i3, max3 := findMax(data)
	delete(data, i3)

	fmt.Printf("max = { index: %d, value: %d }\n", i1, max1)
	fmt.Printf("max = { index: %d, value: %d }\n", i2, max2)
	fmt.Printf("max = { index: %d, value: %d }\n", i3, max3)

	fmt.Printf("total = %d\n", max1+max2+max3)
}

func findMax(data map[int][]int) (int, int) {
	idx := -1
	max := -1
	for i, bag := range data {
		s := sum(bag)
		if s > max {
			max = s
			idx = i
		}
	}
	return idx, max
}

func sum(arr []int) int {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	return sum
}

func readInput() (map[int][]int, error) {
	f, err := os.Open("input.dat")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	result := make(map[int][]int)
	i := 0
	for scanner.Scan() {
		if result[i] == nil {
			result[i] = make([]int, 0)
		}
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			i++
		} else {
			num, err := strconv.Atoi(line)
			if err != nil {
				return nil, err
			}
			result[i] = append(result[i], num)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
