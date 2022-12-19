package main

func (g *Graph) CalcPressure(limit int) int {
	return g.calc([]*Valve{g.start}, 0, 0, limit)
}

func (g *Graph) calc(visited []*Valve, pressure, rate, limit int) int {
	res := pressure + limit*rate // pressure released by waiting
	if limit < 2 {
		return res
	}
	last := visited[len(visited)-1]
	for _, v := range g.operatingNodes {
		if !contains(visited, v) {
			time := g.distances[last][v] + 1 // +1 for opening valve
			if time < limit {
				res = max(res, g.calc(append(visited, v),
					pressure+time*rate,
					rate+v.Rate,
					limit-time,
				))
			}
		}
	}
	return res
}

func contains(arr []*Valve, v *Valve) bool {
	for _, item := range arr {
		if v == item {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
