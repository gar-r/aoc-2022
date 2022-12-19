package main

type Sensor struct {
	Location      *Point
	ClosestBeacon *Point
	Radius        int
	Perimeter     map[int]*Interval
}

func NewSensor(x, y, bx, by int) *Sensor {
	s := &Sensor{
		Location:      &Point{X: x, Y: y},
		ClosestBeacon: &Point{X: bx, Y: by},
	}
	s.Radius = s.Location.Distance(s.ClosestBeacon)
	s.calcPerimeter()
	return s
}

func (s *Sensor) calcPerimeter() {
	north := s.North()
	east := s.East()
	south := s.South()
	west := s.West()
	s.Perimeter = make(map[int]*Interval)
	for p := north; !p.Equal(east); p = p.Offset(1, 1) {
		s.setPerimeter(p)
	}
	for p := east; !p.Equal(south); p = p.Offset(-1, 1) {
		s.setPerimeter(p)
	}
	for p := south; !p.Equal(west); p = p.Offset(-1, -1) {
		if p.Equal(south) {
			s.setPerimeter(p)
		} else {
			s.updatePerimeter(p)
		}
	}
	for p := west; !p.Equal(north); p = p.Offset(1, -1) {
		s.updatePerimeter(p)
	}
}

func (s *Sensor) setPerimeter(p *Point) {
	s.Perimeter[p.Y] = NewInterval(p.X, p.X)
}

func (s *Sensor) updatePerimeter(p *Point) {
	iv, ok := s.Perimeter[p.Y]
	if ok {
		iv.Lower = p.X
	}
}

func (s *Sensor) North() *Point {
	return s.Location.Offset(0, -s.Radius)
}

func (s *Sensor) East() *Point {
	return s.Location.Offset(s.Radius, 0)
}

func (s *Sensor) South() *Point {
	return s.Location.Offset(0, s.Radius)
}

func (s *Sensor) West() *Point {
	return s.Location.Offset(-s.Radius, 0)
}
