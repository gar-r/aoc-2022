package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirection_Turn(t *testing.T) {

	t.Run("Turn Left", func(t *testing.T) {
		assert.Equal(t, Left, Up.Turn(Left))
		assert.Equal(t, Up, Right.Turn(Left))
		assert.Equal(t, Right, Down.Turn(Left))
		assert.Equal(t, Down, Left.Turn(Left))
	})

	t.Run("Turn Right", func(t *testing.T) {
		assert.Equal(t, Right, Up.Turn(Right))
		assert.Equal(t, Down, Right.Turn(Right))
		assert.Equal(t, Left, Down.Turn(Right))
		assert.Equal(t, Up, Left.Turn(Right))
	})

	t.Run("invalid Direction", func(t *testing.T) {
		assert.Equal(t, Up, Up.Turn(Up))
		assert.Equal(t, Down, Down.Turn(Down))
	})

}
