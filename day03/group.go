package main

type group struct {
	rucksacks []*rucksack
}

func newGroup(s1, s2, s3 string) *group {
	g := &group{
		rucksacks: make([]*rucksack, 3),
	}
	g.rucksacks[0] = newRucksack(s1)
	g.rucksacks[1] = newRucksack(s2)
	g.rucksacks[2] = newRucksack(s3)
	return g
}

func (g *group) findBadge() (item, bool) {
	dups := make([]item, 0)
	for _, i := range g.rucksacks[0].items {
		if g.rucksacks[1].contains(i) {
			dups = append(dups, i)
		}
	}
	for _, i := range dups {
		if g.rucksacks[2].contains(i) {
			return i, true
		}
	}
	return '0', false
}
