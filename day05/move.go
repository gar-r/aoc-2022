package main

import (
	"regexp"
	"strconv"
)

type move struct {
	num  int
	from int
	to   int
}

var moveRe = regexp.MustCompile(`move (\d*) from (\d*) to (\d*)`)

func newMove(s string) *move {
	match := moveRe.FindStringSubmatch(s)
	if len(match) != 4 {
		panic(s)
	}
	return &move{
		num:  mustParseInt(match[1]),
		from: mustParseInt(match[2]) - 1,
		to:   mustParseInt(match[3]) - 1,
	}
}

func mustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
