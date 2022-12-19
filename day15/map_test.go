package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMap(t *testing.T) {

	m := NewMap([]*Sensor{
		NewSensor(0, 0, 0, 4),
		NewSensor(-5, -5, -8, -5),
		NewSensor(5, 5, 10, 5),
	})

	t.Run("test borders", func(t *testing.T) {
		tl := NewPoint(-8, -8)
		br := NewPoint(10, 10)
		assert.Equal(t, tl, m.TopLeft)
		assert.Equal(t, br, m.BottomRight)
	})

}
