package main

type rope struct {
	knots   []*pos
	tailLog map[pos]int
}

func newRope(nKnots int) *rope {
	return &rope{
		knots:   initKnots(nKnots),
		tailLog: make(map[pos]int),
	}
}

func initKnots(nKnots int) []*pos {
	arr := make([]*pos, nKnots)
	for i := 0; i < nKnots; i++ {
		arr[i] = &pos{}
	}
	return arr
}

func (r *rope) drag(dir direction) {
	head := r.knots[0]
	head.move(dir)
	r.propagate(r.knots[1:], head)
}

func (r *rope) propagate(knots []*pos, head *pos) {
	if len(knots) == 0 {
		return
	}
	k := knots[0]
	f := force(head, k)
	if !f.null() {
		k.move(f)
		r.propagate(knots[1:], k)
	}
	if len(knots) == 1 {
		r.logTailPos(k)
	}
}

func (r *rope) logTailPos(k *pos) {
	r.tailLog[*k]++
}

func force(p, q *pos) direction {
	dx := p.x - q.x
	dy := p.y - q.y
	if abs(dx) > 1 || abs(dy) > 1 {
		return direction{normalize(dx), normalize(dy)}
	}
	return direction{0, 0}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func normalize(f int) int {
	if f == 0 {
		return 0
	}
	if f < 0 {
		return -1
	}
	return 1
}
