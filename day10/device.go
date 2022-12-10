package main

type Device struct {
	c        *Cpu
	r        *Register
	s        *Crt
	Measures map[int]int
}

func NewDevice() *Device {
	return &Device{
		c: &Cpu{},
		r: &Register{1},
		s: &Crt{make([]string, 240)},
	}
}

func (d *Device) ExecProgram(program []Instruction) {
	for _, instr := range program {
		d.c.LoadInstruction(instr)
		for {
			done := d.c.Tick()
			d.s.Draw(d.c.CycleCount, d.r.Value)
			if d.ShouldMeasure() {
				d.Measures[d.c.CycleCount] = d.SignalStrength()
			}
			if done {
				instr.Effect()(d.r)
				break
			}
		}
	}
}

func (d *Device) SignalStrength() int {
	return d.c.CycleCount * d.r.Value
}

func (d *Device) ShouldMeasure() bool {
	_, ok := d.Measures[d.c.CycleCount]
	return ok
}

type Register struct {
	Value int
}

type Cpu struct {
	CycleCount  int
	instr       Instruction
	instrCycles int
}

func (c *Cpu) LoadInstruction(instr Instruction) {
	c.instr = instr
	c.instrCycles = instr.Cost()
}

func (c *Cpu) Tick() bool {
	c.CycleCount++
	c.instrCycles--
	if c.instrCycles == 0 {
		return true
	}
	return false
}
