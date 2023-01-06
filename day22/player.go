package main

type Player struct {
	row, col int
	dir      Direction
}

func (pp *Player) Code() int {
	return (pp.row+1)*1000 + (pp.col+1)*4 + int(pp.dir)
}

func (pp *Player) Route(b Board, p Path) {
	for _, step := range p {
		pp.route(b, step)
	}
}

func (pp *Player) route(b Board, step Instruction) {
	switch step.(type) {
	case Move:
		pp.move(b, step.(Move))
	case Turn:
		pp.turn(step.(Turn))
	default:
		panic("invalid instruction")
	}
}

func (pp *Player) turn(turn Turn) {
	pp.dir = pp.dir.Turn(turn.Dir)
}

func (pp *Player) move(b Board, move Move) {
	for i := 0; i < move.Amount; i++ {
		nx, ny := b.Next(pp.row, pp.col, pp.dir)
		tile := b[nx][ny]
		if tile == Wall {
			break
		}
		if tile == Empty {
			panic("invalid tile")
		}
		pp.row = nx
		pp.col = ny
	}
}
