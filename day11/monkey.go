package main

type MonkeyInspectFn func(i int64) int64
type MonkeyTargetingFn func(i int64) int64

type Monkey struct {
	Items       []int64
	InspectFn   MonkeyInspectFn
	TargetingFn MonkeyTargetingFn
}

func (m *Monkey) InspectItem() {
	m.Items[0] = m.InspectFn(m.Items[0])
}

func (m *Monkey) FinishInspect() {
	m.Items[0] = m.Items[0] % D
}

func (m *Monkey) ThrowItem(others []*Monkey) {
	item := m.Items[0]
	m.Items = m.Items[1:]
	target := m.TargetingFn(item)
	others[target].CatchItem(item)
}

func (m *Monkey) CatchItem(item int64) {
	m.Items = append(m.Items, item)
}

func ModuloTargetingFn(div, ifTrue, ifFalse int64) MonkeyTargetingFn {
	return func(i int64) int64 {
		if i%div == 0 {
			return ifTrue
		}
		return ifFalse
	}
}

func MulInspectingFn(mul int64) MonkeyInspectFn {
	return func(i int64) int64 {
		return i * mul
	}
}

func SquareInspectingFn() MonkeyInspectFn {
	return func(i int64) int64 {
		return i * i
	}
}

func AddInspectingFn(add int64) MonkeyInspectFn {
	return func(i int64) int64 {
		return i + add
	}
}
