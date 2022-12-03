package main

type rucksack struct {
	items        []item
	compartments []*compartment
}

type compartment struct {
	items   []item
	itemMap map[item]int
}

type item rune

func newRucksack(s string) *rucksack {
	compLen := len(s) / 2
	return &rucksack{
		items: []item(s),
		compartments: []*compartment{
			newCompartment(s[0:compLen]),
			newCompartment(s[compLen:]),
		},
	}
}

func newCompartment(s string) *compartment {
	items := make([]item, len(s))
	itemMap := make(map[item]int)
	for i, c := range s {
		it := item(c)
		items[i] = it
		itemMap[it]++
	}
	return &compartment{
		items:   items,
		itemMap: itemMap,
	}
}

func (r *rucksack) findDuplicate() (item, bool) {
	for _, it := range r.compartments[0].items {
		if r.compartments[1].contains(it) {
			return it, true
		}
	}
	return '0', false
}

func (r *rucksack) contains(it item) bool {
	for _, i := range r.items {
		if i == it {
			return true
		}
	}
	return false
}

func (c *compartment) contains(it item) bool {
	_, ok := c.itemMap[it]
	return ok
}

func (i item) priority() int {
	return priorityMap[rune(i)]
}
