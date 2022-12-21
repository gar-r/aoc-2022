package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGale_Next(t *testing.T) {
	g := parseGale("<>><")
	assert.Equal(t, Left, g.Next())
	assert.Equal(t, Right, g.Next())
	assert.Equal(t, Right, g.Next())
	assert.Equal(t, Left, g.Next())
	assert.Equal(t, Left, g.Next())
	assert.Equal(t, Right, g.Next())
}
