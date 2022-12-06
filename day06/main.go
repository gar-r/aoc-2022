package main

import "fmt"

func main() {
	startPacket := findMarkerIdx(input, 4) + 1
	startMessage := findMarkerIdx(input, 14) + 1

	fmt.Printf("packet: %d\n", startPacket)
	fmt.Printf("message: %d\n", startMessage)
}

func findMarkerIdx(s string, l int) int {
	if len(s) < l {
		return -1
	}
	for i := l - 1; i < len(s); i++ {
		chars := make(map[rune]int)
		for j := i; j > i-l; j-- {
			r := rune(s[j])
			chars[r] = chars[r] + 1
		}
		ok := true
		for _, n := range chars {
			if n > 1 {
				ok = false
			}
		}
		if ok {
			return i
		}
	}
	return -1
}
