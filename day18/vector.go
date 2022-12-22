package main

import "fmt"

type Vector struct {
	X, Y, Z int
}

func (v *Vector) Add(w *Vector) *Vector {
	return &Vector{
		X: v.X + w.X,
		Y: v.Y + w.Y,
		Z: v.Z + w.Z,
	}
}

func (v *Vector) Scale(f int) *Vector {
	return &Vector{
		X: v.X * f,
		Y: v.Y * f,
		Z: v.Z * f,
	}
}

func (v *Vector) Equal(w *Vector) bool {
	return v.X == w.X && v.Y == w.Y && v.Z == w.Z
}

func (v *Vector) String() string {
	return fmt.Sprintf("(%d,%d,%d)", v.X, v.Y, v.Z)
}

var X = &Vector{1, 0, 0}
var Y = &Vector{0, 1, 0}
var Z = &Vector{0, 0, 1}

var MX = X.Scale(-1)
var MY = Y.Scale(-1)
var MZ = Z.Scale(-1)

var Directions = []*Vector{X, Y, Z, MX, MY, MZ}
