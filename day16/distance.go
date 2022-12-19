package main

import "math"

func (g *Graph) ShortestDistance(from *Valve) map[*Valve]int {
	visited := make(map[*Valve]bool)
	distance := make(map[*Valve]int)
	nodes := make([]*Valve, 0)
	for _, v := range g.nodes {
		nodes = append(nodes, v)
		visited[v] = false
		distance[v] = math.MaxInt
	}
	current := from
	distance[from] = 0
	for len(nodes) > 0 && current != nil {
		neighbors := current.Neighbors(g)
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				if distance[neighbor] > distance[current]+1 {
					distance[neighbor] = distance[current] + 1
				}
			}
		}
		visited[current] = true
		nodes = remove(nodes, current)
		current = next(distance, visited)
	}
	return distance
}

func remove(arr []*Valve, item *Valve) []*Valve {
	res := make([]*Valve, 0)
	for _, v := range arr {
		if v == item {
			continue
		}
		res = append(res, item)
	}
	return res
}

func next(distance map[*Valve]int, visited map[*Valve]bool) *Valve {
	var res *Valve
	min := math.MaxInt
	for node, visited := range visited {
		if !visited {
			dst := distance[node]
			if dst < min {
				min = dst
				res = node
			}
		}
	}
	return res
}
