package main

import (
	"regexp"
	"strings"
)

type container struct {
	stacks []*stack
}

type stack struct {
	data []string
}

func (s *stack) pop() string {
	item := s.data[0]
	s.data = s.data[1:]
	return item
}

func (s *stack) push(i string) {
	s.data = append([]string{i}, s.data...)
}

var contRe = regexp.MustCompile(`[\[\s](.)[\]\s]\s?`)

func newContainer(data []string) *container {
	var cont *container
	lastRowIdx := len(data) - 1
	for rowIdx := lastRowIdx; rowIdx >= 0; rowIdx-- {
		line := data[rowIdx]
		matches := contRe.FindAllStringSubmatch(line, len(line))
		if rowIdx == lastRowIdx {
			cont = emptyContainer(len(matches))
		}
		for i, match := range matches {
			s := match[1]
			if strings.TrimSpace(s) == "" {
				continue
			}
			cont.stacks[i].push(s)
		}
	}
	return cont
}

func emptyContainer(n int) *container {
	cont := &container{
		stacks: make([]*stack, n),
	}
	for i := 0; i < n; i++ {
		cont.stacks[i] = &stack{}
	}
	return cont
}
