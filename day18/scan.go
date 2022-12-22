package main

type Scan struct {
	part *Partition
}

func NewScan(data []*Vector) *Scan {
	s := &Scan{
		part: NewPartition(),
	}
	for _, v := range data {
		s.part.Set(v)
	}
	return s
}

func (s *Scan) Contains(v *Vector) bool {
	return s.part.Contains(v)
}
