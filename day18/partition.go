package main

type Partition struct {
	data map[int]map[int]map[int]*Vector
}

func NewPartition() *Partition {
	return &Partition{
		data: make(map[int]map[int]map[int]*Vector),
	}
}

func (p *Partition) Set(v *Vector) {
	px, ok := p.data[v.X]
	if !ok {
		px = make(map[int]map[int]*Vector)
		p.data[v.X] = px
	}
	py, ok := px[v.Y]
	if !ok {
		py = make(map[int]*Vector)
		px[v.Y] = py
	}
	py[v.Z] = v
}

func (p *Partition) Contains(v *Vector) bool {
	px, ok := p.data[v.X]
	if !ok {
		return false
	}
	py, ok := px[v.Y]
	if !ok {
		return false
	}
	_, ok = py[v.Z]
	return ok
}
