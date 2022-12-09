package main

type Stack struct {
	data []*Directory
}

func (s *Stack) Pop() *Directory {
	if len(s.data) == 0 {
		return nil
	}
	idx := len(s.data) - 1
	item := s.data[idx]
	s.data = s.data[:idx]
	return item
}

func (s *Stack) Peek() *Directory {
	if len(s.data) == 0 {
		return nil
	}
	idx := len(s.data) - 1
	return s.data[idx]
}

func (s *Stack) Push(d *Directory) {
	s.data = append(s.data, d)
}
