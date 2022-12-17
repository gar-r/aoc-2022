package cave

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePos(t *testing.T) {

	t.Run("parse position", func(t *testing.T) {
		p := ParsePos("123,456")
		assert.Equal(t, 123, p.W)
		assert.Equal(t, 456, p.D)
	})

	t.Run("parse error", func(t *testing.T) {

		tests := []string{
			"foo",
			"1234",
			"",
		}

		for _, test := range tests {
			t.Run(test, func(t *testing.T) {
				p := ParsePos(test)
				assert.Equal(t, 0, p.W)
				assert.Equal(t, 0, p.D)
			})
		}
	})

}

func TestPos_Trace(t *testing.T) {

	t.Run("490,23 -> 492,23", func(t *testing.T) {
		p1 := &Pos{490, 23}
		p2 := &Pos{492, 23}
		pl := p1.Trace(p2)
		assert.Len(t, pl, 3)
		assertPos(t, pl[0], 490, 23)
		assertPos(t, pl[1], 491, 23)
		assertPos(t, pl[2], 492, 23)
	})

	t.Run("510,23 -> 508,23", func(t *testing.T) {
		p1 := &Pos{510, 23}
		p2 := &Pos{508, 23}
		pl := p1.Trace(p2)
		assert.Len(t, pl, 3)
		assertPos(t, pl[0], 510, 23)
		assertPos(t, pl[1], 509, 23)
		assertPos(t, pl[2], 508, 23)
	})

	t.Run("492,20 -> 492,22", func(t *testing.T) {
		p1 := &Pos{492, 20}
		p2 := &Pos{492, 22}
		pl := p1.Trace(p2)
		assert.Len(t, pl, 3)
		assertPos(t, pl[0], 492, 20)
		assertPos(t, pl[1], 492, 21)
		assertPos(t, pl[2], 492, 22)
	})

	t.Run("492,18 -> 492,16", func(t *testing.T) {
		p1 := &Pos{492, 18}
		p2 := &Pos{492, 16}
		pl := p1.Trace(p2)
		assert.Len(t, pl, 3)
		assertPos(t, pl[0], 492, 18)
		assertPos(t, pl[1], 492, 17)
		assertPos(t, pl[2], 492, 16)
	})

}
