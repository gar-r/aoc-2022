package main

import (
	"fmt"
	"math"
	"strconv"
)

func RunOnAllPartitions(arr []*Valve, fn func(p1, p2 []*Valve)) {
	n := int(math.Pow(2, float64(len(arr)-1)))
	for i := 1; i < n; i++ {
		percent := float64(i) / float64(n) * 100
		fmt.Printf("%f%%\n", percent)
		p1, p2 := split(arr, i)
		fn(p1, p2)
	}
}

func split(arr []*Valve, pn int) (p1, p2 []*Valve) {
	p1 = make([]*Valve, 0)
	p2 = make([]*Valve, 0)
	for i, c := range binPattern(len(arr), pn) {
		if c == '0' {
			p2 = append(p2, arr[i])
		} else {
			p1 = append(p1, arr[i])
		}
	}
	return
}

func binPattern(l, pn int) string {
	bin := strconv.FormatInt(int64(pn), 2)
	format := fmt.Sprintf("%%0%ds", l)
	return fmt.Sprintf(format, bin)
}
