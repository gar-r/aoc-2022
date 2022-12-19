package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInterval_Subtract(t *testing.T) {

	t.Run("i1 contains i2", func(t *testing.T) {
		i1 := NewInterval(0, 10)
		i2 := NewInterval(3, 7)
		res := i1.Subtract(i2)
		assert.Len(t, res, 2)
		assertInterval(t, res[0], 0, 2)
		assertInterval(t, res[1], 8, 10)
	})

	t.Run("i2 overlaps i1 upper segment", func(t *testing.T) {
		i1 := NewInterval(0, 10)
		i2 := NewInterval(5, 15)
		res := i1.Subtract(i2)
		assert.Len(t, res, 1)
		assertInterval(t, res[0], 0, 4)
	})

	t.Run("i2 overlaps i1 lower segment", func(t *testing.T) {
		i1 := NewInterval(0, 10)
		i2 := NewInterval(-5, 7)
		res := i1.Subtract(i2)
		assert.Len(t, res, 1)
		assertInterval(t, res[0], 8, 10)
	})

	t.Run("i2 overlaps i1", func(t *testing.T) {
		i1 := NewInterval(0, 10)
		i2 := NewInterval(-5, 15)
		res := i1.Subtract(i2)
		assert.Len(t, res, 0)
	})

}

func TestInterval_Enumerate(t *testing.T) {
	v := &Interval{Lower: -3, Upper: 1}
	actual := v.Enumerate()
	expected := []int{-3, -2, -1, 0, 1}
	assert.Equal(t, expected, actual)
}
