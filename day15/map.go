package main

import "math"

type Map struct {
	Sensors     []*Sensor
	TopLeft     *Point
	BottomRight *Point
}

func NewMap(sensors []*Sensor) *Map {
	m := &Map{
		Sensors:     sensors,
		TopLeft:     NewPoint(math.MaxInt, math.MaxInt),
		BottomRight: NewPoint(math.MinInt, math.MinInt),
	}
	m.calcBorders()
	return m
}

func (m *Map) FindUncovered(y int) []*Interval {
	uncovered := m.FullRange()
	for _, s := range m.Sensors {
		p, ok := s.Perimeter[y]
		if ok {
			uncovered = SubtractAll(uncovered, p)
			if len(uncovered) == 0 {
				break
			}
		}
	}
	return uncovered
}

func (m *Map) FullRange() []*Interval {
	return []*Interval{NewInterval(m.TopLeft.X, m.BottomRight.X)}
}

func (m *Map) IsSensor(p *Point) bool {
	for _, s := range m.Sensors {
		if p.Equal(s.Location) {
			return true
		}
	}
	return false
}

func (m *Map) IsBeacon(p *Point) bool {
	for _, s := range m.Sensors {
		if p.Equal(s.ClosestBeacon) {
			return true
		}
	}
	return false
}

func (m *Map) calcBorders() {
	for _, s := range m.Sensors {
		if s.West().X < m.TopLeft.X {
			m.TopLeft.X = s.West().X
		}
		if s.East().X > m.BottomRight.X {
			m.BottomRight.X = s.East().X
		}
		if s.North().Y < m.TopLeft.Y {
			m.TopLeft.Y = s.North().Y
		}
		if s.South().Y > m.BottomRight.Y {
			m.BottomRight.Y = s.South().Y
		}
	}
}
