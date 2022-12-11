package main

type Stats struct {
	Inspects map[*Monkey]int
}

func NewStats() *Stats {
	return &Stats{
		Inspects: make(map[*Monkey]int),
	}
}

func (s *Stats) Record(m *Monkey) {
	s.Inspects[m]++
}

func (s *Stats) Merge(other *Stats) {
	for k, v := range other.Inspects {
		s.Inspects[k] += v
	}
}

func (s *Stats) InspectValues() []int {
	res := make([]int, 0)
	for _, v := range s.Inspects {
		res = append(res, v)
	}
	return res
}
