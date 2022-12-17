package cave

func (c *Cave) RunSim(source *Pos) int {
	units := 0
	for {
		s := c.SimSand(source)
		if s == nil {
			return units
		}
		units++
	}
}

func (c *Cave) RunSim2(source *Pos) int {
	units := 0
	for {
		s := c.SimSand(source)
		units++
		if source.W == s.W && source.D == s.D {
			break
		}
	}
	return units
}

func (c *Cave) SimSand(source *Pos) *Pos {
	cp := source
	for {
		np := c.simStep(cp)
		if np == nil {
			c.Set(cp, Sand)
			break
		}
		if np.D >= c.Depth() {
			return nil // cannot rest
		}
		cp = np
	}
	return cp
}

func (c *Cave) simStep(source *Pos) *Pos {
	for _, dir := range []Direction{Down, LeftDown, RightDown} {
		pos := source.Offset(dir)
		if c.blocked(pos) {
			continue
		}
		return pos
	}
	return nil
}

func (c *Cave) blocked(pos *Pos) bool {
	if pos.D >= c.Depth() {
		return false
	}
	tile := c.Get(pos)
	return tile == Brick || tile == Sand
}
