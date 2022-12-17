package cave

import "strings"

type Path struct {
	junctions []*Pos
}

func ParsePath(s string) (path *Path) {
	parts := strings.Split(s, " -> ")
	path = &Path{junctions: make([]*Pos, len(parts))}
	for i, part := range parts {
		path.junctions[i] = ParsePos(part)
	}
	return
}

func (p *Path) Trace() []*Pos {
	res := make([]*Pos, 0)
	if len(p.junctions) == 1 {
		res = append(res, p.junctions[0])
		return res
	}
	for i := 1; i < len(p.junctions); i++ {
		j1 := p.junctions[i-1]
		j2 := p.junctions[i]
		res = append(res, j1.Trace(j2)...)
	}
	return res
}
