package main

type pos struct {
	x, y int
}

func (p *pos) move(dir direction) {
	p.x += dir.dx
	p.y += dir.dy
}

func (p *pos) clone() *pos {
	return &pos{
		x: p.x,
		y: p.y,
	}
}
