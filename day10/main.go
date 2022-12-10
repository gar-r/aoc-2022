package main

import "fmt"

func main() {
	prog1 := parse(example)
	prog2 := parse(input)

	fmt.Printf("program 1 line count: %d\n", len(prog1))
	fmt.Printf("program 2 line count: %d\n", len(prog2))

	dev := NewDevice()
	dev.Measures = map[int]int{
		20:  0,
		60:  0,
		100: 0,
		140: 0,
		180: 0,
		220: 0,
	}
	dev.ExecProgram(prog2)

	total := 0
	for _, v := range dev.Measures {
		total += v
	}

	fmt.Printf("total signal strength: %d\n", total)

	fmt.Println()
	dev.s.Print()

}
