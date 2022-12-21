package main

type RockOrder struct {
	order []Rock
	step  int
}

func (r *RockOrder) Next() Rock {
	rock := r.order[r.step]
	r.step++
	if r.step == len(r.order) {
		r.step = 0
	}
	return rock
}

type Rock interface {
	// BlockPattern lists rock offsets calculated from bottom-left corner
	BlockPattern() []*Offset
}

// H is a horizontal 4x1 line of blocks
type H struct{}

func (h *H) BlockPattern() []*Offset {
	return []*Offset{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
	}
}

// V is a vertical 1x4 line of blocks
type V struct{}

func (v *V) BlockPattern() []*Offset {
	return []*Offset{
		{0, 0},
		{0, 1},
		{0, 2},
		{0, 3},
	}
}

// P is a "+" shaped rock
type P struct{}

func (p *P) BlockPattern() []*Offset {
	return []*Offset{
		{0, 1},
		{1, 0},
		{1, 1},
		{1, 2},
		{2, 1},
	}
}

// L is a reverse "L" shaped rock
type L struct{}

func (l *L) BlockPattern() []*Offset {
	return []*Offset{
		{0, 0},
		{1, 0},
		{2, 0},
		{2, 1},
		{2, 2},
	}
}

// B is a 2x2 box shaped rock
type B struct{}

func (b *B) BlockPattern() []*Offset {
	return []*Offset{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}
}

type Offset struct {
	Dx, Dy int
}

func (o *Offset) Apply(x, y int) (int, int) {
	return x + o.Dx, y + o.Dy
}
