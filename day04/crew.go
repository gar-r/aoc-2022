package main

import "strings"

type crew struct {
	zones []*zone
}

func newCrew(s string) *crew {
	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		panic(s)
	}
	zones := make([]*zone, 2)
	zones[0] = newZone(parts[0])
	zones[1] = newZone(parts[1])
	return &crew{
		zones: zones,
	}
}

func (c *crew) isRedundant() bool {
	return c.zones[0].contains(c.zones[1]) || c.zones[1].contains(c.zones[0])
}

func (c *crew) hasOverlaps() bool {
	return c.zones[0].overlaps(c.zones[1])
}
