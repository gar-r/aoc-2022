package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Round struct {
	opponent     Figure
	desiredState State
}

func readInput() ([]*Round, error) {
	f, err := os.Open("input.dat")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	result := make([]*Round, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		moves := strings.Split(scanner.Text(), " ")
		if len(moves) != 2 {
			return nil, fmt.Errorf("invalid row: %v", moves)
		}
		figure, err := ParseFigure(moves[0])
		if err != nil {
			return nil, err
		}
		state, err := ParseState(moves[1])
		if err != nil {
			return nil, err
		}
		result = append(result, &Round{
			opponent:     figure,
			desiredState: state,
		})
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
