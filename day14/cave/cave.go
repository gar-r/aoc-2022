package cave

import "strings"

type Cave struct {
	data [][]Tile
}

func NewCave(w, d int) *Cave {
	c := &Cave{
		data: make([][]Tile, w),
	}
	for i := 0; i < w; i++ {
		c.data[i] = make([]Tile, d)
	}
	return c
}

func (c *Cave) Get(p *Pos) Tile {
	return c.data[p.W][p.D]
}

func (c *Cave) Set(p *Pos, v Tile) {
	c.data[p.W][p.D] = v
}

func (c *Cave) Width() int {
	return len(c.data)
}

func (c *Cave) Depth() int {
	return len(c.data[0])
}

func (c *Cave) ActualDepth() int {
	for d := c.Depth() - 1; d >= 0; d-- {
		for w := 0; w < c.Width(); w++ {
			if c.Get(&Pos{W: w, D: d}) == Brick {
				return d
			}
		}
	}
	return 0
}

func Parse(s string, w, d int) *Cave {
	pl := parsePathList(s)
	cave := NewCave(w, d)
	for _, p := range pl {
		for _, pos := range p.Trace() {
			cave.Set(pos, Brick)
		}
	}
	return cave
}

func parsePathList(s string) []*Path {
	lines := strings.Split(s, "\n")
	pl := make([]*Path, len(lines))
	for i, line := range lines {
		pl[i] = ParsePath(line)
	}
	return pl
}

func (c *Cave) String() string {
	sb := &strings.Builder{}
	for d := 0; d < c.Depth(); d++ {
		for w := 450; w < 575; w++ {
			sb.WriteString(c.data[w][d].String())
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
