package main

import (
	"fmt"
	"math"
)

func main() {
	data := input
	scan := NewScan(data)
	size := cubeSize(data)
	tracer := NewTracer(scan, size)
	tracer.Flood = true
	surfaces := tracer.TraceSurfaces()
	fmt.Printf("surfaces: %d\n", surfaces)
}

// a relatively crude method to find a bounding cube for the input
func cubeSize(arr []*Vector) int {
	minC := math.MaxInt
	maxC := math.MinInt
	for _, v := range arr {
		if v.X < minC {
			minC = v.X
		}
		if v.X > maxC {
			maxC = v.X
		}
		if v.Y < minC {
			minC = v.Y
		}
		if v.Y > maxC {
			maxC = v.Y
		}
		if v.Z < minC {
			minC = v.Z
		}
		if v.Z > maxC {
			maxC = v.Z
		}
	}
	return max(abs(minC), abs(maxC)) + 3
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
