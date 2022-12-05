package main

import (
	"fmt"
	"strconv"
	"strings"
)

type zone struct {
	startId int
	endId   int
}

func newZone(s string) *zone {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		panic(fmt.Sprintf("invalid zone: %s", s))
	}
	i1, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(s)
	}
	i2, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(s)
	}
	return &zone{
		startId: i1,
		endId:   i2,
	}
}

func (z *zone) contains(other *zone) bool {
	return z.startId <= other.startId && z.endId >= other.endId
}

func (z *zone) overlaps(other *zone) bool {
	return z.endId <= other.endId && z.endId >= other.startId ||
		z.startId >= other.startId && z.startId <= other.endId
}
