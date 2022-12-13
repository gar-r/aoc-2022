package main

type HeightMap struct {
	Start *Coord
	Goal  *Coord
	data  [][]int32
}

func NewHeightMap(input []string) *HeightMap {
	hm := &HeightMap{}
	hm.data = make([][]int32, 0)
	for i, r := range input {
		row := make([]int32, 0)
		for j, c := range r {
			if rune(c) == 'S' {
				hm.Start = &Coord{Row: i, Col: j}
			}
			if rune(c) == 'E' {
				hm.Goal = &Coord{Row: i, Col: j}
			}
			row = append(row, c)
		}
		hm.data = append(hm.data, row)
	}
	return hm
}

func (h *HeightMap) Elevation(c *Coord) int32 {
	e := h.data[c.Row][c.Col]
	if rune(e) == 'S' {
		return 'a'
	}
	if rune(e) == 'E' {
		return 'z'
	}
	return e
}

func (h *HeightMap) Neighbors(c *Coord) map[Direction]*Coord {
	res := make(map[Direction]*Coord, 0)
	for _, d := range []Direction{North, East, South, West} {
		if h.hasNeighbor(c, d) {
			res[d] = c.Offset(d)
		}
	}
	return res
}

func (h *HeightMap) hasNeighbor(c *Coord, d Direction) bool {
	p := c.Offset(d)
	return p.Row >= 0 &&
		p.Row < len(h.data) &&
		p.Col >= 0 &&
		p.Col < len(h.data[0])
}

func (h *HeightMap) canClimb(origin, target *Coord) bool {
	diff := h.Elevation(target) - h.Elevation(origin)
	return diff <= 1
}

func (h *HeightMap) canDescend(origin, target *Coord) bool {
	return h.canClimb(target, origin)
}
