package cave

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsePath(t *testing.T) {
	s := "488,148 -> 488,144 -> 488,148 -> 490,148"
	j := ParsePath(s).junctions
	assert.Len(t, j, 4)
	assertPos(t, j[0], 488, 148)
	assertPos(t, j[1], 488, 144)
	assertPos(t, j[2], 488, 148)
	assertPos(t, j[3], 490, 148)
}

func TestPath_Trace(t *testing.T) {
	s := "488,148 -> 488,144 -> 488,148 -> 490,148"
	pl := ParsePath(s).Trace()
	assert.Len(t, pl, 13)
	assertPos(t, pl[0], 488, 148)
	assertPos(t, pl[1], 488, 147)
	assertPos(t, pl[2], 488, 146)
	assertPos(t, pl[3], 488, 145)
	assertPos(t, pl[4], 488, 144)
	assertPos(t, pl[5], 488, 144)
	assertPos(t, pl[6], 488, 145)
	assertPos(t, pl[7], 488, 146)
	assertPos(t, pl[8], 488, 147)
	assertPos(t, pl[9], 488, 148)
	assertPos(t, pl[10], 488, 148)
	assertPos(t, pl[11], 489, 148)
	assertPos(t, pl[12], 490, 148)
}

func assertPos(t *testing.T, p *Pos, w, d int) {
	t.Helper()
	assert.Equal(t, p.W, w)
	assert.Equal(t, p.D, d)
}
