package main

import "fmt"

type Interval struct {
	Lower int
	Upper int
}

func NewInterval(lower, upper int) *Interval {
	return &Interval{lower, upper}
}

func (v *Interval) Contains(n int) bool {
	return v.Lower <= n && n <= v.Upper
}

func (v *Interval) Subtract(other *Interval) []*Interval {
	if v.Contains(other.Lower) {
		if v.Contains(other.Upper) {
			return []*Interval{
				NewInterval(v.Lower, other.Lower-1),
				NewInterval(other.Upper+1, v.Upper),
			}
		}
		return []*Interval{NewInterval(v.Lower, other.Lower-1)}
	}
	if v.Contains(other.Upper) {
		return []*Interval{NewInterval(other.Upper+1, v.Upper)}
	}
	if other.Contains(v.Lower) && other.Contains(v.Upper) {
		return []*Interval{}
	}
	return []*Interval{v}
}

func (v *Interval) Enumerate() []int {
	res := make([]int, 0)
	for i := v.Lower; i <= v.Upper; i++ {
		res = append(res, i)
	}
	return res
}

func (v *Interval) String() string {
	return fmt.Sprintf("(%d,%d)", v.Lower, v.Upper)
}

func SubtractAll(arr []*Interval, v *Interval) []*Interval {
	res := make([]*Interval, 0)
	for _, i := range arr {
		res = append(res, i.Subtract(v)...)
	}
	return res
}
