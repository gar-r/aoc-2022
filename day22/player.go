package main

type Player struct {
	row, col int
	dir      Direction
	cube     bool
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
		if pp.cube {
			pp.moveCube(b, step.(Move))
		} else {
			pp.move(b, step.(Move))
		}
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
	pp.moveGeneric(b, move, func(x, y int, dir Direction) (int, int, Direction) {
		nx, ny := b.Next(x, y, dir)
		return nx, ny, dir
	})
}

func (pp *Player) moveCube(b Board, move Move) {
	pp.moveGeneric(b, move, func(x, y int, dir Direction) (int, int, Direction) {
		return b.Wrap(x, y, dir)
	})
}

func (pp *Player) moveGeneric(b Board, move Move, nextFn NextFunc) {
	for i := 0; i < move.Amount; i++ {
		nx, ny, nd := nextFn(pp.row, pp.col, pp.dir)
		tile := b[nx][ny]
		if tile == Wall {
			break
		}
		if tile == Empty {
			panic("invalid tile")
		}
		pp.row = nx
		pp.col = ny
		pp.dir = nd
	}
}

type NextFunc func(int, int, Direction) (int, int, Direction)
