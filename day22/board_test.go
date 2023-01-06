package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoard_Next(t *testing.T) {

	b := parseBoard(`
  ..#
  .#.
...#..
...#.#
`)

	tests := []struct {
		x, y   int
		dir    Direction
		nx, ny int
	}{
		{0, 0, Right, 0, 2},
		{0, 1, Right, 0, 2},
		{0, 2, Right, 0, 3},
		{0, 4, Right, 0, 2},
		{0, 0, Left, 0, 4},
		{0, 1, Left, 0, 4},
		{0, 2, Left, 0, 4},
		{2, 0, Up, 3, 0},
		{2, 1, Up, 3, 1},
		{2, 2, Up, 1, 2},
		{3, 0, Down, 2, 0},
		{3, 1, Down, 2, 1},
		{3, 2, Down, 0, 2},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			nx, ny := b.Next(test.x, test.y, test.dir)
			assert.Equal(t, test.nx, nx)
			assert.Equal(t, test.ny, ny)
		})
	}

}
