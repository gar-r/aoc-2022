package main

import "fmt"

type Graph struct {
	nodes          map[string]*Valve
	operatingNodes map[string]*Valve
	distances      DistanceMap
	start          *Valve
}
type DistanceMap map[*Valve]map[*Valve]int

func NewGraph(valves []*Valve, start string) *Graph {
	g := &Graph{
		nodes:          make(map[string]*Valve),
		operatingNodes: make(map[string]*Valve),
		distances:      make(DistanceMap),
	}
	for _, v := range valves {
		g.nodes[v.Id] = v
		if v.Rate > 0 {
			g.operatingNodes[v.Id] = v
		}
	}
	g.start = g.nodes[start]
	for _, v := range g.nodes {
		g.distances[v] = g.ShortestDistance(v)
	}
	return g
}

type Valve struct {
	Id      string
	Rate    int
	Tunnels []string
}

func (v *Valve) Neighbors(g *Graph) []*Valve {
	res := make([]*Valve, 0)
	for _, t := range v.Tunnels {
		n, ok := g.nodes[t]
		if ok {
			res = append(res, n)
		}
	}
	return res
}

func (v *Valve) String() string {
	return fmt.Sprintf("{id: %s, rate: %d} => %v",
		v.Id, v.Rate, v.Tunnels)
}
