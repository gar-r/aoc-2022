package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSensor(t *testing.T) {

	s := NewSensor(0, 0, 0, 4)

	t.Run("radius", func(t *testing.T) {
		assert.Equal(t, 4, s.Radius)
	})

	t.Run("perimeter", func(t *testing.T) {
		p := s.Perimeter
		assertInterval(t, p[0], -4, 4)
		assertInterval(t, p[1], -3, 3)
		assertInterval(t, p[2], -2, 2)
		assertInterval(t, p[3], -1, 1)
		assertInterval(t, p[4], 0, 0)
		assertInterval(t, p[-1], -3, 3)
		assertInterval(t, p[-2], -2, 2)
		assertInterval(t, p[-3], -1, 1)
		assertInterval(t, p[-4], 0, 0)
	})

}

func assertInterval(t *testing.T, iv *Interval, lower, upper int) {
	t.Helper()
	assert.Equal(t, iv.Lower, lower)
	assert.Equal(t, iv.Upper, upper)
}
