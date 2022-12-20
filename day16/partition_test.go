package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartition(t *testing.T) {

	v1 := &Valve{}
	v2 := &Valve{}
	v3 := &Valve{}
	v4 := &Valve{}

	arr := []*Valve{v1, v2, v3, v4}

	p1, p2 := split(arr, 7)

	assert.Equal(t, []*Valve{v2, v3, v4}, p1)
	assert.Equal(t, []*Valve{v1}, p2)

}
