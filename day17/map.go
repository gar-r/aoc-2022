package main

const ChamberWidth = 7
const DefaultExpand = 10

type Map struct {
	order *RockOrder
	gale  *Gale
	data  [][]bool
}

func NewMap(input string) *Map {
	m := &Map{
		order: &RockOrder{
			order: []Rock{&H{}, &P{}, &L{}, &V{}, &B{}},
		},
		gale: parseGale(input),
	}
	m.expand(DefaultExpand)
	return m
}

func (m *Map) SimRock() {
	rock := m.order.Next()
	_, y := m.Peak()
	m.autoExpand()
	x := 2
	y = y + 3 + 1
	for {
		// uncomment the following line to draw map to terminal
		//m.render(rock, x, y)
		x = m.applyGale(rock, x, y)
		if m.isColliding(rock, x, y-1) {
			break
		}
		y--
	}
	m.settle(rock, x, y)
}

func (m *Map) HasBlock(x, y int) bool {
	return m.data[y][x]
}

func (m *Map) SetBlock(x, y int) {
	m.data[y][x] = true
}

func (m *Map) autoExpand() {
	_, y := m.Peak()
	if len(m.data)-y <= 8 {
		m.expand(DefaultExpand)
	}
}

func (m *Map) expand(n int) {
	for i := 0; i < n; i++ {
		m.data = append(m.data, make([]bool, 7))
	}
}

func (m *Map) Peak() (x, y int) {
	for y := len(m.data) - 1; y >= 0; y-- {
		for x, b := range m.data[y] {
			if b {
				return x, y
			}
		}
	}
	return 0, -1 // no block
}

func (m *Map) applyGale(rock Rock, x, y int) int {
	nx := m.gale.Next().Apply(x)
	if !m.isColliding(rock, nx, y) {
		x = nx
	}
	return x
}

func (m *Map) isColliding(rock Rock, x, y int) bool {
	for _, block := range rock.BlockPattern() {
		nx, ny := block.Apply(x, y)
		if m.isOutOfBounds(nx, ny) || m.HasBlock(nx, ny) {
			return true
		}
	}
	return false
}

func (m *Map) isOutOfBounds(x, y int) bool {
	return x < 0 || y < 0 || x >= ChamberWidth
}

func (m *Map) settle(rock Rock, x, y int) {
	for _, block := range rock.BlockPattern() {
		nx, ny := block.Apply(x, y)
		m.SetBlock(nx, ny)
	}
}
