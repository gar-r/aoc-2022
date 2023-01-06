package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePath(t *testing.T) {
	actual := parsePath("10R5L5R10")
	expected := Path{
		Move{10},
		Turn{Right},
		Move{5},
		Turn{Left},
		Move{5},
		Turn{Right},
		Move{10},
	}
	assert.Equal(t, expected, actual)
}

func TestParseBoard(t *testing.T) {
	actual := parseBoard(`
  ..#
  .#.
...#..
...#.#
`)

	expected := Board{
		{Empty, Empty, Free, Free, Wall, Empty},
		{Empty, Empty, Free, Wall, Free, Empty},
		{Free, Free, Free, Wall, Free, Free},
		{Free, Free, Free, Wall, Free, Wall},
	}

	assert.Equal(t, expected, actual)

}
