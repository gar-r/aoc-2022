package packet

type Stack struct {
	data []Packet
}

func (s *Stack) Pop() Packet {
	if len(s.data) == 0 {
		return nil
	}
	idx := len(s.data) - 1
	item := s.data[idx]
	s.data = s.data[:idx]
	return item
}

func (s *Stack) Push(p Packet) {
	s.data = append(s.data, p)
}

func (s *Stack) Empty() bool {
	return len(s.data) == 0
}
