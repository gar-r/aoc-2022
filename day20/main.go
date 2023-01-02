package main

import "fmt"

func main() {
	part1()
	part2()
}

func part2() {
	f := NewFile(input)
	f.ApplyKey(811589153)
	for i := 0; i < 10; i++ {
		f.Mix()
	}
	fmt.Printf("groove coordinates: %d\n", calcGrove(f))
}

func part1() {
	f := NewFile(input)
	f.Mix()
	fmt.Printf("groove coordinates: %d\n", calcGrove(f))
}

func calcGrove(f *File) int {
	start := f.Zero()
	x := f.seek(start, 1000)
	y := f.seek(x, 1000)
	z := f.seek(y, 1000)
	return x.Value + y.Value + z.Value
}
