package main

type Instruction interface {
	Cost() int
	Effect() func(*Register)
}

type Addx struct {
	Value int
}

func (a *Addx) Cost() int {
	return 2
}

func (a *Addx) Effect() func(*Register) {
	return func(r *Register) {
		r.Value += a.Value
	}
}

type Noop struct{}

func (n *Noop) Cost() int {
	return 1
}

func (n *Noop) Effect() func(*Register) {
	return func(*Register) {}
}
