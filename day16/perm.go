package main

func Perms(
	valves []*Valve,
	limit func(p []*Valve) bool,
	exec func(p []*Valve),
) {
	p(valves, limit, exec, nil)
}

func p(
	valves []*Valve,
	limit func(p []*Valve) bool,
	exec func(p []*Valve),
	res []*Valve,
) {
	if limit(res) {
		return // !
	}
	if len(valves) == 0 {
		exec(res)
	}
	for _, v := range valves {
		p(del(valves, v), limit, exec, append(res, v))
	}
}

func del(arr []*Valve, item *Valve) []*Valve {
	res := make([]*Valve, 0)
	for _, v := range arr {
		if v == item {
			continue
		}
		res = append(res, v)
	}
	return res
}
