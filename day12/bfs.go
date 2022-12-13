package main

func (h *HeightMap) FindShortestPath() int {
	steps := 0
	q := &Queue{}
	q.Enqueue(h.Start)
	visited := make(map[Coord]bool)
	visited[*h.Start] = true
	for !q.Empty() {
		ql := q.Len()
		for i := 0; i < ql; i++ {
			item := q.Dequeue()
			if item.Equal(h.Goal) {
				return steps
			}
			neighbors := h.Neighbors(item)
			for _, n := range neighbors {
				_, v := visited[*n]
				if !v && h.canClimb(item, n) {
					visited[*n] = true
					q.Enqueue(n)
				}
			}
		}
		steps++
	}
	panic("path not found")
}

func (h *HeightMap) FindShortestPathToAny() int {
	steps := 0
	q := &Queue{}
	q.Enqueue(h.Goal)
	visited := make(map[Coord]bool)
	visited[*h.Goal] = true
	for !q.Empty() {
		ql := q.Len()
		for i := 0; i < ql; i++ {
			item := q.Dequeue()
			if h.Elevation(item) == 'a' {
				return steps
			}
			neighbors := h.Neighbors(item)
			for _, n := range neighbors {
				_, v := visited[*n]
				if !v && h.canDescend(item, n) {
					visited[*n] = true
					q.Enqueue(n)
				}
			}
		}
		steps++
	}
	panic("path not found")
}
