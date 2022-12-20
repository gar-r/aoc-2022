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
		if !contains(visited, v.Id) {
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

func (g *Graph) ElephantSim(limit int) int {
	res := 0
	RunOnAllPartitions(nodeList(g.operatingNodes), func(p1, p2 []*Valve) {
		res1 := g.calc(append(p2, g.start), 0, 0, limit)
		res2 := g.calc(append(p1, g.start), 0, 0, limit)
		res = max(res, res1+res2)
	})
	return res
}

func nodeList(m map[string]*Valve) []*Valve {
	res := make([]*Valve, 0)
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func contains(arr []*Valve, id string) bool {
	for _, item := range arr {
		if id == item.Id {
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
