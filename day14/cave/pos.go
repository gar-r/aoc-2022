package cave

import (
	"strconv"
	"strings"
)

type Pos struct {
	W, D int
}

func (p *Pos) Offset(dir Direction) *Pos {
	switch dir {
	case LeftDown:
		return &Pos{W: p.W - 1, D: p.D + 1}
	case RightDown:
		return &Pos{W: p.W + 1, D: p.D + 1}
	default:
		return &Pos{W: p.W, D: p.D + 1}
	}
}

func (p *Pos) Trace(t *Pos) []*Pos {
	res := make([]*Pos, 0)
	if p.W == t.W {
		if p.D < t.D {
			for i := p.D; i <= t.D; i++ {
				res = append(res, &Pos{p.W, i})
			}
		} else {
			for i := p.D; i >= t.D; i-- {
				res = append(res, &Pos{p.W, i})
			}
		}
	} else {
		if p.W < t.W {
			for i := p.W; i <= t.W; i++ {
				res = append(res, &Pos{i, p.D})
			}
		} else {
			for i := p.W; i >= t.W; i-- {
				res = append(res, &Pos{i, p.D})
			}
		}
	}
	return res
}

func ParsePos(s string) (p *Pos) {
	parts := strings.Split(s, ",")
	p = &Pos{}
	if len(parts) == 2 {
		p.W, _ = strconv.Atoi(parts[0])
		p.D, _ = strconv.Atoi(parts[1])
	}
	return
}
